use regex::Regex;
use std::{collections::HashMap, io::Read};

/*
Traverse the nodes starting With AAA and ending with ZZZ, following the directions alternately
*/
fn traverse(
    nodes: &HashMap<String, Vec<String>>,
    start_node: Option<&str>,
    directions: Vec<char>,
) -> i32 {
    let mut total_steps = 0;
    let mut node_to_check = match start_node {
        Some(node) => node,
        None => "AAA",
    };

    let mut directions_index = 0;

    while !node_to_check.contains("Z") {
        let next_direction = match directions[directions_index % directions.len()] {
            'L' => 0,
            'R' => 1,
            _ => panic!("Invalid direction"),
        };

        println!("Checking node {}", node_to_check);
        let next_node = match &nodes.get(node_to_check) {
            Some(node) => node[next_direction].as_str(),
            None => panic!("Node not found"),
        };

        node_to_check = next_node;
        directions_index += 1;
        total_steps += 1;
    }

    return total_steps;
}

fn main() {
    let args = std::env::args().collect::<Vec<String>>();

    if args.len() != 2 {
        println!("Usage: {} <number>", args[0]);
        return;
    }

    let file_path = &args[1];
    let mut file = std::fs::File::open(file_path).unwrap();

    // read the file into a string
    let contents = &mut String::new();
    file.read_to_string(contents).unwrap();

    let lines: Vec<&str> = contents.lines().collect();

    let directions = lines[0].chars().collect::<Vec<char>>();

    let mut nodes: HashMap<String, Vec<String>> = HashMap::new();
    let mut iterations = 0;
    println!("START INDEXING");
    for line in lines[2..].iter() {
        let re = Regex::new(r"(\w{3}) = \((\w{3}), (\w{3})\)").unwrap();
        let caps = re.captures(line).unwrap();

        let actual = caps.get(1).unwrap().as_str();
        let left = caps.get(2).unwrap().as_str();
        let right = caps.get(3).unwrap().as_str();

        iterations += 1;
        if iterations % 50 == 0 {
            println!("{} iterations", iterations);
        }

        nodes.insert(
            actual.to_string(),
            vec![left.to_string(), right.to_string()],
        );
    }

    println!("START TRAVERSING");

    // part 1
    // let result = traverse(&nodes, None, directions.clone());
    // println!("Result: {}", result);

    // part 2
    let all_node_with_a = nodes
        .iter()
        .filter(|(key, _)| key.contains('A'))
        .map(|(key, _)| key.to_string())
        .collect::<Vec<String>>();

    all_node_with_a.iter().for_each(|node| {
        let result = traverse(&nodes, Some(node), directions.clone());
        println!("Result: {}", result);
    });
}
