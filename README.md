# ft_linear_regression

An introduction to machine learning — predicting car prices from mileage using linear regression trained with gradient descent.

This is a 42 school project where the goal is to implement a simple linear regression algorithm from scratch, without relying on any machine learning library.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [How It Works](#how-it-works)
  - [Linear Regression](#linear-regression)
  - [Cost Function (MSE)](#cost-function-mse)
  - [Gradient Descent](#gradient-descent)
- [Usage](#usage)
- [Visualisation](#visualisation)

## Overview

The program predicts the **price** of a car based on its **mileage (km)** using a linear hypothesis function trained with gradient descent. The dataset contains 24 data points mapping kilometres to price.

The project is split into three commands:

| Command | Description |
|---------|-------------|
| `train` | Reads the dataset, runs gradient descent, saves the learned parameters |
| `predict` | Prompts for a mileage and outputs the estimated price |
| `precision` | Computes the MSE and relative error of the model |

## Project Structure

```
ft_linear_regression/
├── cmd/
│   ├── train/          # Training program — runs gradient descent
│   │   └── main.go
│   ├── predict/        # Prediction program — estimates price for a given km
│   │   └── main.go
│   └── precision/      # Bonus — computes MSE and precision of the model
│       └── main.go
├── model/
│   ├── dataset.go      # DataSet and Row types, MSE computation, CSV conversion
│   └── parameters.go   # Parameters (theta0, theta1) and the hypothesis function
├── training/
│   ├── gradient.go     # Gradient descent algorithm with convergence check
│   └── cost.go         # Cost computation with feature normalisation + denormalisation
├── csvio/
│   └── csvio.go        # CSV reading/writing for dataset and parameters
├── consts/
│   └── consts.go       # File path constants
├── plot/
│   └── plot.go         # Scatter plot + regression line visualisation (gonum/plot)
├── data/
│   ├── data.csv        # Training dataset (km, price)
│   └── parameters.csv  # Learned parameters (theta0, theta1)
├── images/
│   └── scatter.png     # Output plot after training
├── go.mod
└── go.sum
```

## How It Works

### Linear Regression

Linear regression fits a straight line through the data to model the relationship between an input variable **x** (mileage) and an output variable **y** (price).

The **hypothesis function** is:

$$h_\theta(x) = \theta_0 + \theta_1 \cdot x$$

Where:
- $\theta_0$ is the **y-intercept** (bias) — the predicted price when mileage is 0
- $\theta_1$ is the **slope** (weight) — how much the price changes per unit of mileage
- $x$ is the input feature (mileage in km)

The goal is to find the values of $\theta_0$ and $\theta_1$ that best fit the data.

![Linear regression example](https://upload.wikimedia.org/wikipedia/commons/thumb/b/be/Normdist_regression.png/325px-Normdist_regression.png)

*A linear regression line fitted through a set of data points.*

### Cost Function (MSE)

To measure how well our line fits the data, we use a **cost function**. The **Mean Squared Error (MSE)** quantifies the average squared difference between predictions and actual values:

$$J(\theta_0, \theta_1) = \frac{1}{m} \sum_{i=0}^{m-1} \left( h_\theta(x_i) - y_i \right)^2$$

Where:
- $m$ is the number of training examples (24 in our dataset)
- $h_\theta(x_i)$ is the predicted price for the $i$-th car
- $y_i$ is the actual price of the $i$-th car
- $J(\theta_0, \theta_1)$ is the cost — a single number representing how wrong the model is

The lower $J$, the better the fit. The goal of training is to **minimise** $J(\theta_0, \theta_1)$.

When plotted as a function of $\theta_0$ and $\theta_1$, the cost function forms a 3D convex surface (a bowl shape). The minimum of this bowl is where our optimal parameters lie.

![3D cost function](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a3/Gradient_descent_in_linear_regression.svg/800px-Gradient_descent_in_linear_regression.svg.png)

*The cost function $J(\theta_0, \theta_1)$ forms a convex surface. Gradient descent iteratively moves toward the minimum.*

### Gradient Descent

**Gradient descent** is an optimisation algorithm that iteratively adjusts $\theta_0$ and $\theta_1$ to minimise the cost function. At each step, it computes the **partial derivatives** (gradients) of $J$ with respect to each parameter and moves in the direction of steepest descent.

The update rules are:

$$\theta_0 := \theta_0 - \alpha \cdot \frac{1}{m} \sum_{i=0}^{m-1} \left( h_\theta(x_i) - y_i \right)$$

$$\theta_1 := \theta_1 - \alpha \cdot \frac{1}{m} \sum_{i=0}^{m-1} \left( h_\theta(x_i) - y_i \right) \cdot x_i$$

Where:
- $\alpha$ is the **learning rate** — controls the step size (set to `0.01` in this project)
- Both $\theta_0$ and $\theta_1$ are updated **simultaneously** at each iteration
- The algorithm stops when either:
  - The changes become smaller than $\varepsilon = 10^{-9}$ (convergence)
  - The maximum number of iterations (1,000,000) is reached

**Feature normalisation:** Since mileage values are large (22,899 to 240,000), the raw gradient would be enormous and cause the algorithm to diverge. The training normalises the km values to the $[0, 1]$ range using min-max scaling before running gradient descent, then denormalises the resulting parameters so they work with real-world values.

## Usage

```bash
# Train the model
go run cmd/train/main.go

# Predict a price
go run cmd/predict/main.go

# Check model precision (bonus)
go run cmd/precision/main.go
```

Training output:
```
converged at iteration 1789
theta0: 8499.60, theta1: -0.02
```

Prediction example:
```
Input a kilometrage: 100000
Predicted price: 6499.6
```

## Visualisation

After training, a scatter plot with the regression line is saved to `images/scatter.png`.

The plot shows:
- **Blue dots**: the training data (km vs price)
- **Red line**: the learned linear model $h_\theta(x) = \theta_0 + \theta_1 \cdot x$
