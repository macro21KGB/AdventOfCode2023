use std::env;

#[derive(Debug, Clone, PartialEq)]
enum Direction {
    UP,
    DOWN,
    LEFT,
    RIGHT,
    UNKNOWN,
}

#[derive(Debug, Clone, PartialEq)]
struct PipeOpening {
    pointed_to: Direction,
    location_on_pipe: Direction,
}

#[derive(Debug, Clone, PartialEq)]
struct Pipe {
    openings: (Direction, Direction),
    position: (u32, u32),
    distance: Option<u32>,
    raw: char,
}

fn convert_matrix_chart_to_pipes(matrix: Vec<Vec<char>>) -> Vec<Vec<Pipe>> {
    let mut pipes: Vec<Vec<Pipe>> = Vec::new();

    for (row_index, row) in matrix.iter().enumerate() {
        let mut row_of_pipes: Vec<Pipe> = Vec::new();

        for (column_index, pipe_symbol) in row.iter().enumerate() {
            if *pipe_symbol == '.' {
                continue;
            }

            let pipe = Pipe::new(*pipe_symbol, (row_index as u32, column_index as u32));
            row_of_pipes.push(pipe);
        }

        pipes.push(row_of_pipes);
    }

    return pipes;
}

impl Pipe {
    fn new(pipe_symbol: char, position: (u32, u32)) -> Pipe {
        match pipe_symbol {
            '|' => Pipe {
                position,
                openings: (Direction::UP, Direction::DOWN),
                distance: None,
                raw: pipe_symbol,
            },
            '-' => Pipe {
                position,
                openings: (Direction::LEFT, Direction::RIGHT),
                distance: None,
                raw: pipe_symbol,
            },
            '7' => Pipe {
                position,
                openings: (Direction::DOWN, Direction::LEFT),
                distance: None,
                raw: pipe_symbol,
            },
            'L' => Pipe {
                position,
                openings: (Direction::UP, Direction::RIGHT),
                distance: None,
                raw: pipe_symbol,
            },
            'J' => Pipe {
                position,
                openings: (Direction::LEFT, Direction::UP),
                distance: None,
                raw: pipe_symbol,
            },
            'F' => Pipe {
                position,
                openings: (Direction::RIGHT, Direction::DOWN),
                distance: None,
                raw: pipe_symbol,
            },
            'S' => Pipe {
                position,
                openings: (Direction::UNKNOWN, Direction::UNKNOWN),
                distance: None,
                raw: pipe_symbol,
            },
            invalid_symbol => panic!("Invalid pipe symbol: {}", invalid_symbol),
        }
    }
}

fn main() {
    println!("Hello, world!");

    let file_name = env::args().nth(1).expect("No file name provided");

    let file_contents = std::fs::read_to_string(file_name).expect("Could not read file");

    let matrix: Vec<Vec<char>> = file_contents
        .lines()
        .map(|line| line.chars().collect())
        .collect();

    let pipes = convert_matrix_chart_to_pipes(matrix);

    println!("{:?}", pipes);
}
