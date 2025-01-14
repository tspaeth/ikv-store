use std::sync::Arc;

use log::info;

use crate::index::ckv::CKVIndex;
use crate::kafka::consumer::IKVKafkaConsumer;
use crate::kafka::processor::WritesProcessor;
use crate::proto::generated_proto::common::IKVStoreConfig;

use super::index_loader;

#[cfg(test)]
#[path = "main_test.rs"]
mod main_test;

// TODO: change backend url
// const SERVER_URL: &str = "https://gateway-writer-alb-1-1639339774.us-west-2.elb.amazonaws.com:443";

/// Stateful controller for managing IKV core key-val storage.
pub struct Controller {
    index: Arc<CKVIndex>,
    processor: Arc<WritesProcessor>,
    kafka_consumer: IKVKafkaConsumer,
}

impl Controller {
    pub fn open(config: &IKVStoreConfig) -> anyhow::Result<Self> {
        // fetch server configs and override|merge with client supplied configs
        // let config = Controller::merge_with_server_config(client_supplied_config)?;

        // Load index
        index_loader::load_index(&config)?;
        let index = Arc::new(CKVIndex::open_or_create(&config)?);

        // Initialize kafka consumer
        let processor = Arc::new(WritesProcessor::new(index.clone()));
        let kafka_consumer = IKVKafkaConsumer::new(&config, processor.clone())?;

        // Start write event consumption
        // Blocks till pending events are consumed
        // Consumes incoming events in background thereafter

        // TODO: kafka consumer code needs to be reviewed!!
        kafka_consumer.run_in_background()?;

        Ok(Controller {
            index,
            processor,
            kafka_consumer,
        })
    }

    /// Atomic reference to the index.
    pub fn index_ref(&self) -> Arc<CKVIndex> {
        self.index.clone()
    }

    /// Atomic reference to the writes processor.
    pub fn writes_processor_ref(&self) -> Arc<WritesProcessor> {
        self.processor.clone()
    }

    pub fn close(self) -> anyhow::Result<()> {
        self.kafka_consumer.stop();
        info!("Closing IKV, Bye Bye.");
        Ok(())
    }

    /*
    /// Unused, there are issues with using client TLS certificate
    /// Client code (ex. java reader) will fetch and merge with
    /// server supplied config temporarily.
    fn merge_with_server_config(
        client_supplied_config: &IKVStoreConfig,
    ) -> anyhow::Result<IKVStoreConfig> {
        let mut config = Self::fetch_server_configs(client_supplied_config)?;

        // override with client_supplied_config
        for (k, v) in client_supplied_config.stringConfigs.iter() {
            config.stringConfigs.insert(k.to_string(), v.to_string());
        }
        for (k, v) in client_supplied_config.intConfigs.iter() {
            config.intConfigs.insert(k.to_string(), *v);
        }
        for (k, v) in client_supplied_config.floatConfigs.iter() {
            config.floatConfigs.insert(k.to_string(), *v);
        }
        for (k, v) in client_supplied_config.bytesConfigs.iter() {
            config.bytesConfigs.insert(k.to_string(), v.clone());
        }
        for (k, v) in client_supplied_config.booleanConfigs.iter() {
            config.booleanConfigs.insert(k.to_string(), *v);
        }

        Ok(config)
    }

    #[tokio::main(flavor = "current_thread")]
    pub async fn fetch_server_configs(
        client_supplied_config: &IKVStoreConfig,
    ) -> anyhow::Result<IKVStoreConfig> {
        // Build request
        let account_id = client_supplied_config
            .stringConfigs
            .get("account_id")
            .ok_or(anyhow!("account_id is a required config"))?;

        let account_passkey = client_supplied_config
            .stringConfigs
            .get("account_passkey")
            .ok_or(anyhow!("account_passkey is a required config"))?;

        let store_name = client_supplied_config
            .stringConfigs
            .get("store_name")
            .ok_or(anyhow!("store_name is a required config"))?;

        let request = tonic::Request::new(GetUserStoreConfigRequest {
            user_store_context_initializer: Some(UserStoreContextInitializer {
                credentials: Some(AccountCredentials {
                    account_id: account_id.to_string(),
                    account_passkey: account_passkey.to_string(),
                }),
                store_name: store_name.to_string(),
            }),
        });

        let tls = ClientTlsConfig::new().domain_name("www.inlined.io");

        let channel = Channel::from_static(SERVER_URL)
            .tls_config(tls)?
            .connect()
            .await?;

        let mut client = InlineKvWriteServiceClient::new(channel);

        let tonic_response = client.get_user_store_config(request).await?;

        let response = tonic_response.get_ref();
        debug!("Fetched server-side configs: = {:?}", response);

        if response.global_config.is_none() {
            return Ok(IKVStoreConfig::default());
        }

        let server_config = response.global_config.clone().unwrap();

        // TODO: resolve multiple proto objects from protoc and tonic
        let mut copied_config = IKVStoreConfig::new();
        copied_config.stringConfigs = server_config.string_configs;
        copied_config.intConfigs = server_config.int_configs;
        copied_config.floatConfigs = server_config.float_configs;
        copied_config.bytesConfigs = server_config.bytes_configs;
        copied_config.booleanConfigs = server_config.boolean_configs;

        Ok(copied_config)
    }
    */
}
