// enum RSEnum {
//     Foo(i32),
//     Bar(String),
//     Baz(Vec<String>),
// }

fn main() {
    let foo = Some(5);

    if let Some(value) = foo {

    }

    match foo {
        Some(value) => println!("{}", value),
        None => println!("None"),
    }

    // Closure, like JS's .map.
    foo.map(|x| {
    });

    foo.filter(|x| *x < 10);

    // match foo {
    //     RSEnum::Foo(value) => println!("Foo: {}", value),
    //     RSEnum::Bar(value) => println!("Bar: {}", value),
    //     RSEnum::Baz(value) => println!("Baz: {:?}", value),
    // }

}
