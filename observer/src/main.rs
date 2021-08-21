use std::borrow::Cow;

use quick_protobuf::{deserialize_from_slice, serialize_into_vec};

mod model;

fn main() {
    println!("Hello, world!");
    let a = model::auth::PrivateKey { X: Cow::from("hello".to_string()), Y: Cow::from("hello".to_string()) };
    let s = serialize_into_vec(&a).unwrap();
    println!("{:?}", &s);
    let b = deserialize_from_slice::<model::auth::PrivateKey>(&s).unwrap();
    println!("{:?}", &b)
}
