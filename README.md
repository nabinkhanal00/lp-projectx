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

You can then visit to `localhost:3000` to interact with the calculator.

## Run tests
```bash
# inside the calculator directory
# run all tests
go test ./...
```

## Tasks
- [x] Build Simple UI
- [x] Support Addition, Subtraction, Multiplication and Division
- [ ] Support Exponential and Logarithms
- [ ] Support Trigonometric Operations
- [ ] Improve UI/UX

## Comment
I will be happy getting pull requests and issues regarding bug-fixes, improvements and features.
