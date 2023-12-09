use std::collections::HashMap;

#[derive(Debug)]
struct Readings {
    levels: HashMap<i32, Vec<i32>>,
}

fn calculate_difference(first_number: i32, second_number: i32) -> i32 {
    second_number - first_number
}

impl Readings {
    fn calculate_all_layers(&mut self) {
        // calculate the next layer until all values are zeros

        while !self
            .levels
            .get(&(self.levels.len() as i32))
            .unwrap()
            .iter()
            .all(|&x| x == 0)
        {
            self.calculate_next_layer_step();
        }
    }

    fn get_last_number_from_first_layer(&self) -> i32 {
        let first_layer = self.levels.get(&1).unwrap();

        first_layer.last().unwrap().clone()
    }

    fn calculate_next_layer_step(&mut self) {
        let current_layer_index = self.levels.len() as i32;
        let current_layer = self.levels.get(&current_layer_index).unwrap();

        let mut next_layer = Vec::new();

        for i in 0..current_layer.len() - 1 {
            let first_number = current_layer[i];
            let second_number = current_layer[i + 1];

            let difference = calculate_difference(first_number, second_number);

            next_layer.push(difference);
        }

        self.levels.insert(current_layer_index + 1, next_layer);
    }

    fn generate_new_last_value_for_every_layer(&mut self) {
        let mut current_layer_index = (self.levels.len()) as i32;

        while current_layer_index > 0 {
            let current_layer = self.levels.get(&current_layer_index).clone().unwrap();
            let last_value_from_layer = match current_layer.last() {
                Some(value) => value,
                None => &0,
            };

            // get values from the previous layer

            let previous_layer_value = match self.levels.get(&(current_layer_index + 1)) {
                Some(prev_layer) => prev_layer.last().unwrap(),
                None => &0,
            };

            let new_value = last_value_from_layer + previous_layer_value;

            let mut new_layer = current_layer.clone();
            new_layer.push(new_value);

            self.levels.insert(current_layer_index, new_layer);

            current_layer_index -= 1;
        }

        // start from the last layer...
        // calaculate the new value for the last layer
        // and then go back to the previous layer and calculate the new value, and so on
    }

    fn generate_new_first_value_for_every_layer(&mut self) {
        let mut current_layer_index = (self.levels.len()) as i32;

        while current_layer_index > 0 {
            let current_layer = self.levels.get(&current_layer_index).clone().unwrap();
            let first_value_from_layer = match current_layer.first() {
                Some(value) => value,
                None => &0,
            };

            // get values from the previous layer

            let previous_layer_value = match self.levels.get(&(current_layer_index + 1)) {
                Some(prev_layer) => prev_layer.first().unwrap(),
                None => &0,
            };

            let new_value = first_value_from_layer - previous_layer_value;

            let mut new_layer = current_layer.clone();
            new_layer.insert(0, new_value);

            self.levels.insert(current_layer_index, new_layer);

            current_layer_index -= 1;
        }
    }
}

fn parse_readings(readings: &str) -> Readings {
    let values = readings
        .split_whitespace()
        .map(|num| num.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    let mut levels = HashMap::new();
    levels.insert(1, values.clone());
    Readings { levels }
}

fn main() {
    let args = std::env::args().collect::<Vec<String>>();

    if args.len() < 2 {
        panic!("Please provide a file name");
    }

    let filename = &args[1];

    let mut readings = std::fs::read_to_string(filename)
        .expect("Something went wrong reading the file")
        .lines()
        .map(|line| parse_readings(line))
        .collect::<Vec<Readings>>();

    for reading in &mut readings {
        reading.calculate_all_layers();
    }

    for reading in &mut readings {
        reading.generate_new_last_value_for_every_layer();
        reading.generate_new_first_value_for_every_layer();
    }

    let summed_readings = readings
        .iter()
        .map(|reading| {
            let last_value = reading.get_last_number_from_first_layer();
            last_value
        })
        .sum::<i32>();

    println!("Summed readings: {}", summed_readings);

    let summed_readings_first_value = readings
        .iter()
        .map(|reading| {
            let first_value = reading.levels.get(&1).unwrap()[0];
            first_value
        })
        .sum::<i32>();

    println!(
        "Summed readings first value: {}",
        summed_readings_first_value
    );
}

#[cfg(test)]
mod tests {
    #[test]
    fn test_calculate_difference() {
        let first_number = 1;
        let second_number = 2;

        let difference = super::calculate_difference(first_number, second_number);

        assert_eq!(difference, 1);
    }

    #[test]
    fn test_parse_readings() {
        let readings = "1 2 3 4 5 6 7 8 9 10";

        let parsed_readings = super::parse_readings(readings);

        assert_eq!(parsed_readings.levels.len(), 1);
        assert_eq!(parsed_readings.levels.get(&1).unwrap()[0], 1);
        assert_eq!(parsed_readings.levels.get(&1).unwrap().len(), 10);
    }

    #[test]
    fn test_simple_calculation() {
        let readings = "0 3 6 9 12 15";

        let mut parsed_readings = super::parse_readings(readings);

        parsed_readings.calculate_all_layers();
        parsed_readings.generate_new_last_value_for_every_layer();

        assert_eq!(parsed_readings.levels.len(), 3);
        assert_eq!(parsed_readings.levels.get(&3).unwrap()[0], 0);
        assert_eq!(parsed_readings.get_last_number_from_first_layer(), 18);
    }

    #[test]
    fn test_complex_calculation_2() {
        let readings = "10 13 16 21 30 45";

        let mut parsed_readings = super::parse_readings(readings);

        parsed_readings.calculate_all_layers();

        parsed_readings.generate_new_first_value_for_every_layer();
        parsed_readings.generate_new_last_value_for_every_layer();

        for level in 1..=parsed_readings.levels.len() as i32 {
            println!(
                "Level {}: {:?}",
                level,
                parsed_readings.levels.get(&level).unwrap()
            );
        }
    }
}
