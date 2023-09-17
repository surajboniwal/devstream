pub mod user;

use tokio_postgres::{connect, Error, NoTls};

use tonic::transport::Server;
use user::{repo::UserRepository, service::UserService};

use user_rpc::user_service_rpc_server::UserServiceRpcServer;

pub mod user_rpc {
    tonic::include_proto!("user_rpc");
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    let (client, _) = connect("postgres://surajboniwal:devstream@db:5432/devstream", NoTls).await?;

    let user_repository = UserRepository::new(client);

    let user_service: UserService = UserService::new(user_repository);

    let addr = "127.0.0.1:9001".parse().unwrap();

    Server::builder()
        .add_service(UserServiceRpcServer::new(user_service))
        .serve(addr)
        .await;

    Ok(())
}
