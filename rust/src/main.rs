fn main() {
    let res = get_input()
        .lines()
        .map(parse_line)
        .fold(Point {x: 0, y: 0}, |mut acc, p| {
            acc.x += p.x;
            acc.y += p.y;
            return acc;
        });    

    println!("{:?}", res);
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn parse_line(line: &'static str) -> Point { // we can also say &str.
    let (dir, ammount) = line.split_once(" ").expect("must contain a whitespace");
    let ammount = ammount.parse::<i32>().expect("must be an integer");
    // OR let ammount: i32 = ammount.parse().expect("must be an integer");
    // OR let ammount: i32 = str::parse(amount).expect("must be an integer");
    // There are more...
    if dir == "forward" {
        return Point { x: ammount, y: 0 };
    } else if dir == "down" {
        return Point { x: 0, y: ammount };
    } else {
        return Point { x: 0, y: -ammount};
    }
}

fn get_input() -> &'static str {
	return "forward 5
down 5
forward 8
up 3
down 8
forward 2"
}
