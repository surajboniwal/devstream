use sonyflake::Sonyflake;
use tokio_postgres::{Client, Error};

pub struct User {
    pub id: i64,
    pub name: String,
    pub username: String,
    pub email: String,
    pub password: String,
}

pub struct UserRepository {
    client: Client,
    id_generator: Sonyflake,
}

impl UserRepository {
    pub fn new(client: Client, id_generator: Sonyflake) -> Self {
        Self { client, id_generator }
    }

    pub async fn create(
        &self,
        name: String,
        username: String,
        email: String,
        password: String,
    ) -> Result<User, Error> {
        let id: u64 = self.id_generator.next_id().unwrap();

        let result = self.client.execute(
            "INSERT INTO users (id, name, email, username, password) VALUES($1, $2, $3, $4, $5)",
            &[&(id as i64), &name, &email, &username, &password],
        ).await;

        let user = match result {
            Ok(_) => User {
                id: id as i64,
                email,
                name,
                password,
                username,
            },
            Err(err) => return Err(err),
        };

        Ok(user)
    }

    pub async fn get_all(&self) -> Result<Vec<User>, Error> {
        let result = self
            .client
            .query(
                "SELECT id, name, email, username, password from users;",
                &[],
            )
            .await;

        let users = match result {
            Ok(rows) => {
                let mut u: Vec<User> = vec![];

                for row in rows {
                    u.push(User {
                        id: row.get(0),
                        name: row.get(1),
                        email: row.get(2),
                        username: row.get(3),
                        password: row.get(4),
                    });
                }

                u
            }
            Err(err) => return Err(err),
        };

        Ok(users)
    }
}
