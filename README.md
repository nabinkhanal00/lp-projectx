# Powerful Calculator

It is a simple calculator (just called it powerful😁) that provides interface for the user to calculate simple expressions involving:
- Addition
- Subtraction
- Multiplication
- Division 

It supports **decimal numbers**.

## Build Instructions
```bash
# clone the git repository and name it calculator
git clone https://github.com/nabinkhanal00/lp-projectx calculator
cd calculator

# build the docker image
docker build -t calc .

# run the image in a container
docker run -it -p 3000:3000 calc

```

You can then visit to [localhost:3000](http://localhost:3000) to interact with the calculator.

You can also find the app on [https://calc.nabinkhanal00.com.np](https://calc.nabinkhanal00.com.np).

## Run tests
```bash
# inside the calculator directory
# run all tests
go test ./...
```

## Features
1. Perform Addition, Subtraction, Multiplication and Division.
2. Display the steps involved in the calculation.
3. Record all the calculations performed.
4. See the steps and result of all the previous calculations.
5. Delete each history and clear all history.

## Tasks
- [x] Build Simple UI
- [x] Support Addition, Subtraction, Multiplication and Division
- [ ] Support Exponential and Logarithms
- [ ] Support Trigonometric Operations
- [ ] Improve UI/UX

## Comment
I will be happy getting pull requests and issues regarding bug-fixes, improvements, features, and tests.
