pub mod user;

use tokio_postgres::{connect, NoTls};

use tonic::transport::Server;
use user::{repo::UserRepository, service::UserService};

#[tokio::main]
async fn main() {
    let (client, connection) =
        connect("postgres://surajboniwal:devstream@db:5432/devstream", NoTls)
            .await
            .unwrap();

    tokio::spawn(async move {
        if let Err(e) = connection.await {
            eprintln!("connection error: {}", e);
        }
    });

    let user_repository = UserRepository::new(client);

    let user_service = UserService::new(user_repository);

    let addr = "0.0.0.0:3001".parse().unwrap();

    Server::builder()
        .add_service(user_service)
        .serve(addr)
        .await
        .unwrap();
}
