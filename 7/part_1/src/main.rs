use std::collections::HashMap;
use std::fs::read_to_string;

#[derive(Debug, PartialEq, Clone)]
struct Hand<'a> {
    cards: Vec<&'a str>,
    bid: i32,
}

#[derive(Debug, PartialEq, Hash, Eq)]
enum HandType {
    HighCard,
    Pair,
    TwoPairs,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

fn calculate_power(card: &str) -> i32 {
    match card {
        "A" => 14,
        "K" => 13,
        "Q" => 12,
        "J" => 11,
        "T" => 10,
        "9" => 9,
        "8" => 8,
        "7" => 7,
        "6" => 6,
        "5" => 5,
        "4" => 4,
        "3" => 3,
        "2" => 2,
        _ => card.parse().unwrap(),
    }
}

impl<'a> Hand<'a> {
    fn new(cards: Vec<&'a str>, bid: i32) -> Hand<'a> {
        Hand { cards, bid }
    }

    fn is_stronger_than(&self, other: Hand) -> bool {
        if self.cards[0] != other.cards[0] {
            return self.cards[0] < other.cards[0];
        }

        for i in 1..self.cards.len() {
            if self.cards[i] != other.cards[i] {
                return calculate_power(self.cards[i]) > calculate_power(other.cards[i]);
            }
        }

        return false;
    }

    fn get_hand_type(&self) -> HandType {
        let mapped_amounts =
            self.cards
                .iter()
                .fold(HashMap::<&str, i32>::new(), |mut acc, card| {
                    let card_amount = acc.entry(card).or_insert(0);
                    *card_amount += 1;
                    acc
                });

        let pairs = mapped_amounts
            .iter()
            .filter(|(_, amount)| **amount == 2)
            .count();

        let three_of_a_kind = mapped_amounts
            .iter()
            .filter(|(_, amount)| **amount == 3)
            .count();

        let four_of_a_kind = mapped_amounts
            .iter()
            .filter(|(_, amount)| **amount == 4)
            .count();

        let five_of_a_kind = mapped_amounts
            .iter()
            .filter(|(_, amount)| **amount == 5)
            .count();

        if five_of_a_kind == 1 {
            return HandType::FiveOfAKind;
        }

        if pairs == 1 && three_of_a_kind == 1 {
            return HandType::FullHouse;
        }

        if four_of_a_kind == 1 {
            return HandType::FourOfAKind;
        }

        if pairs == 2 {
            return HandType::TwoPairs;
        }

        if three_of_a_kind == 1 {
            return HandType::ThreeOfAKind;
        }

        if pairs == 1 {
            return HandType::Pair;
        }

        return HandType::HighCard;
    }
}

fn find_last_strongest_hand<'a>(hands: Vec<&Hand<'a>>) -> Hand<'a> {
    let mut strongest_hand = hands[0].clone();
    for hand in hands {
        if !hand.is_stronger_than(strongest_hand.clone()) {
            strongest_hand = hand.clone();
        }
    }

    return strongest_hand;
}

fn parse_line(line: &str) -> Hand {
    let splitted_line: Vec<&str> = line.split_whitespace().collect();

    let cards: Vec<&str> = splitted_line[0]
        .split("")
        .filter(|elem| elem.len() != 0)
        .collect();

    let bid: i32 = splitted_line[1].parse().unwrap();

    Hand::new(cards, bid)
}

fn main() {
    // read the argument from the command line
    let args: Vec<String> = std::env::args().collect();

    if args.len() < 2 {
        println!("Please provide a file name");
        return;
    }

    let filename = &args[1];

    // how to read a file in rust
    let contents = read_to_string(filename).unwrap();

    let lines: Vec<&str> = contents.lines().collect();

    let hands = lines
        .iter()
        .map(|line| parse_line(line))
        .collect::<Vec<Hand>>();

    // create a map with HandType as key and all the cards that match that type as value
    let mut mapped_hands_by_type =
        hands
            .iter()
            .fold(HashMap::<HandType, Vec<&Hand>>::new(), |mut acc, hand| {
                let hand_type = hand.get_hand_type();
                let hand_type_vec = acc.entry(hand_type).or_insert(Vec::new());
                hand_type_vec.push(hand);
                acc
            });

    let mut total_sum = 0;
    let mut current_rank = 1;
    // start iterating from the lowest hand type to the highest
    for hand_type in [
        HandType::HighCard,
        HandType::Pair,
        HandType::TwoPairs,
        HandType::ThreeOfAKind,
        HandType::FullHouse,
        HandType::FourOfAKind,
        HandType::FiveOfAKind,
    ] {
        // if there are no hands of this type, continue
        if !mapped_hands_by_type.contains_key(&hand_type) {
            continue;
        }

        // get the hands of this type
        let hands_of_type: &mut Vec<&Hand<'_>> = mapped_hands_by_type.get_mut(&hand_type).unwrap();

        if hands_of_type.len() > 1 {
            while hands_of_type.len() > 0 {
                let lowest_hand = find_last_strongest_hand(hands_of_type.to_vec());
                total_sum += lowest_hand.bid * current_rank;
                println!("{:?}, {}", lowest_hand.bid, current_rank);

                current_rank += 1;
                hands_of_type.retain(|hand| *hand != &lowest_hand);
            }
        }

        if hands_of_type.len() == 1 {
            total_sum += hands_of_type[0].bid * current_rank;
            println!("{:?} {}", hands_of_type[0].bid, current_rank);
            current_rank += 1;
        }
    }

    println!("{}", total_sum);
}
