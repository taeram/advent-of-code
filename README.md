![test workflow](https://github.com/taeram/advent-of-code/actions/workflows/test.yml/badge.svg)

# advent-of-code

Solutions for https://adventofcode.com

## Setup

Run `./bin/setup.sh` to install the local development dependencies.

## Usage

To run a specific day, run: `make run day=2022/day-01`

* This will use the `input.txt` file in specified directory.

To test a specific day, run: `make test day=2022/day-01`

To lint a specific day, run: `make lint day=2022/day-01`

To tidy, lint, test and run a specific day, run: `make ready day=2022/day-01`
