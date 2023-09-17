use postgres::{Client, NoTls, Error};

mod user;

fn main() -> Result<(), Error> {
    let client = Client::connect("postgres://surajboniwal:devstream@db:5432/devstream", NoTls)?;

    let _user_repository: user::UserRepository = user::UserRepository::new(client);

    Ok(())
}
