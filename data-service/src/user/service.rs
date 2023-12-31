use tonic::{Request, Response, Status};
use user_rpc::{user_service_rpc_server::UserServiceRpc, UserRequest, UserResponse};

use self::user_rpc::user_service_rpc_server::UserServiceRpcServer;

use super::repo::UserRepository;

mod user_rpc {
    tonic::include_proto!("user_rpc");
}

pub struct UserService {
    user_repository: UserRepository,
}

impl UserService {
    pub fn new(user_repository: UserRepository) -> UserServiceRpcServer<UserService> {
        UserServiceRpcServer::new(Self { user_repository })
    }
}

#[tonic::async_trait]
impl UserServiceRpc for UserService {
    async fn create(
        &self,
        request: Request<UserRequest>,
    ) -> Result<Response<UserResponse>, Status> {
        let user_request = request.into_inner();

        let result = self
            .user_repository
            .create(
                user_request.name,
                user_request.username,
                user_request.email,
                user_request.password,
            )
            .await;

        match result {
            Ok(user) => Ok(Response::new(UserResponse {
                status: true,
                id: Some(user.id),
                constraint: None,
            })),
            Err(err) => {
                println!("{}", err);
                return Ok(Response::new(UserResponse {
                    status: false,
                    id: None,
                    constraint: Some("error".into()),
                }));
            }
        }
    }
}
