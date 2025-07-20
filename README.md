# Profit Calculator

A modern GUI profit calculator built in Go using the Fyne framework that helps you calculate your business profit, earnings before tax (EBT), and profit ratio with an intuitive graphical interface.

## 📋 About

This profit calculator is a user-friendly GUI application designed to help business owners and entrepreneurs quickly calculate their:
- **EBT (Earnings Before Tax)**: Revenue minus expenses
- **Profit After Tax**: EBT adjusted for tax rate
- **Profit Ratio**: Percentage of profit relative to revenue

The application features a clean, modern interface that takes inputs in cents for precision and provides results in both cents and euros with proper formatting.

## ✨ Features

- **Modern GUI Interface**: Built with Fyne for a native look and feel across platforms
- **Cross-Platform**: Runs on Windows, macOS, and Linux
- **Real-time Calculations**: Instant results as you calculate
- **Input Validation**: Clear error messages for invalid inputs
- **Currency Formatting**: Professional display with comma separators
- **Reset Functionality**: Quick reset button to clear all fields

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or later
- Git (for cloning the repository)
- C compiler (required for Fyne):
  - **Windows**: Install TDM-GCC or similar
  - **macOS**: Install Xcode command line tools: `xcode-select --install`
  - **Linux**: Install development tools: `sudo apt-get install gcc libc6-dev`

### Installation

#### Option 1: Clone and Build from Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/vahiiiid/profit-calculator.git
   cd profit-calculator
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Build the application:**
   ```bash
   go build -o profit-calculator .
   ```

4. **Run the application:**
   ```bash
   ./profit-calculator
   ```

#### Option 2: Direct Go Installation

```bash
go install github.com/vahiiiid/profit-calculator@latest
```

### Usage

1. **Launch the application:**
   ```bash
   ./profit-calculator
   ```
   Or simply double-click the executable file.

2. **Using the interface:**
   - Enter your **Revenue** in cents in the first field (e.g., 10000000 for €100,000.00)
   - Enter your **Expenses** in cents in the second field (e.g., 5000000 for €50,000.00)  
   - Enter your **Tax percentage** with decimal point in the third field (e.g., 19.5 for 19.5%)
   - Click **Calculate** to see your results
   - Click **Reset** to clear all fields and start over

3. **View your results:**
   The application will display:
   - EBT in cents and euros with proper formatting
   - Profit after tax in cents and euros
   - Profit ratio as a percentage

### Example Usage

Using the interface with:
- **Revenue**: 10,000,000 cents (€100,000.00)
- **Expenses**: 5,000,000 cents (€50,000.00)
- **Tax Rate**: 19.5%

**Results:**
```
EBT: 5000000 cents (50,000.00 euros)
Profit: 4025000 cents (40,249.99 euros)
Profit Ratio: 40.25%
```

## 🛠️ Development

### Project Structure

```
profit-calculator/
├── app.go                 # Main GUI application logic
├── calculations.go        # Core calculation functions (testable)
├── calculations_test.go   # Comprehensive test suite
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── LICENSE                # MIT License
├── README.md              # Project documentation
├── screenshot.png         # Application screenshot
└── profit-calculator      # Compiled binary
```

### Dependencies

The application uses the following main dependencies:
- **Fyne v2**: Modern and easy-to-use GUI toolkit for Go
- **Go Standard Library**: For calculations and string formatting

### Building for Different Platforms

```bash
# For Windows
fyne package -os windows -icon icon.png

# For macOS 
fyne package -os darwin -icon icon.png

# For Linux
fyne package -os linux -icon icon.png

# Manual cross-compilation
# For Windows
GOOS=windows GOARCH=amd64 go build -o profit-calculator.exe app.go

# For macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o profit-calculator-macos app.go

# For macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o profit-calculator-macos-arm64 app.go

# For Linux
GOOS=linux GOARCH=amd64 go build -o profit-calculator-linux app.go
```

### Development Setup

1. **Install Fyne command-line tools:**
   ```bash
   go install fyne.io/fyne/v2/cmd/fyne@latest
   ```

2. **Run in development mode:**
   ```bash
   go run .
   ```

3. **Test the build:**
   ```bash
   go build -o profit-calculator .
   ```

## 🧪 Testing

The project includes comprehensive tests for all calculation logic to ensure accuracy and reliability.

### Running Tests

**Run all tests:**
```bash
go test
```

**Run tests with verbose output:**
```bash
go test -v
```

**Run tests with coverage:**
```bash
go test -cover
```

**Run benchmark tests:**
```bash
go test -bench=.
```

**Generate detailed coverage report:**
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Coverage

The test suite covers:

- **Calculation Logic**: Comprehensive tests for profit calculations including:
  - Basic profit calculations with various tax rates
  - Edge cases (zero revenue, zero expenses, losses)
  - Boundary conditions (maximum values, negative values)
  - Different tax rate scenarios (0%, 50%, 100%+)

- **Currency Formatting**: Tests for the euro formatting function including:
  - Positive and negative amounts
  - Large numbers with comma separators
  - Decimal precision and rounding
  - Edge cases and boundary conditions

- **Result Formatting**: Tests for the complete result display format

- **Performance**: Benchmark tests to ensure calculations perform efficiently

### Test Structure

```
profit-calculator/
├── calculations.go      # Core calculation functions (testable)
├── calculations_test.go # Comprehensive test suite
├── app.go              # GUI application (uses calculation functions)
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

The calculation logic has been extracted into separate, testable functions:
- `CalculateProfit()`: Main profit calculation logic
- `FormatEuro()`: Currency formatting with comma separators
- `FormatResult()`: Complete result formatting

### Example Test Cases

The tests include scenarios such as:
- **Basic calculation**: €100,000 revenue, €50,000 expenses, 19.5% tax → €40,250 profit
- **Zero tax**: €10,000 revenue, €6,000 expenses, 0% tax → €4,000 profit  
- **Loss scenario**: €5,000 revenue, €8,000 expenses, 20% tax → €2,400 loss (with tax benefit)
- **Edge cases**: Zero revenue, maximum values, 100% tax rate

## 🤝 Contributing

We welcome contributions to improve the Profit Calculator! Here's how you can help:

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch:**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Run tests and ensure the code builds:**
   ```bash
   go test -v
   go test -cover
   go build -o profit-calculator .
   ```
5. **Test the GUI thoroughly**
6. **Commit your changes:**
   ```bash
   git commit -m 'Add some amazing feature'
   ```
7. **Push to the branch:**
   ```bash
   git push origin feature/amazing-feature
   ```
8. **Open a Pull Request**

### Contribution Guidelines

- Follow Go conventions and best practices
- Follow Fyne UI guidelines for consistency
- Write clear, commented code
- Test your changes thoroughly on different platforms
- Update documentation if necessary
- Be respectful and constructive in discussions

### Ideas for Contributions

- **Enhanced Input Validation**: More robust error handling and user feedback
- **Currency Support**: Support for different currencies beyond euros
- **Calculation History**: Save and review previous calculations
- **Export Features**: Export results to PDF or spreadsheet formats
- **Themes**: Dark/light theme toggle
- **Keyboard Shortcuts**: Improve accessibility
- **GUI Tests**: Automated testing for the user interface
- **Internationalization**: Support for multiple languages
- **Advanced Calculations**: More sophisticated business metrics

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

**MIT License Summary:**
- ✅ Commercial use
- ✅ Modification
- ✅ Distribution
- ✅ Private use
- ❌ Liability
- ❌ Warranty

## 📝 Changelog

### [2.0.0] - 2025-01-XX

#### Added
- **Complete GUI Rewrite**: Transformed from command-line to modern GUI application
- **Fyne Framework Integration**: Cross-platform native GUI experience
- **Real-time Calculations**: Instant feedback with Calculate button
- **Reset Functionality**: Quick reset button to clear all inputs
- **Enhanced Input Validation**: User-friendly error messages
- **Professional Formatting**: Comma-separated number display
- **Cross-platform Support**: Native look and feel on Windows, macOS, and Linux

#### Changed
- **Interface**: Moved from command-line to graphical user interface
- **User Experience**: More intuitive and accessible for all users
- **Installation**: Now requires GUI dependencies but easier to use

#### Removed
- Command-line interface (replaced with GUI)

### [1.0.0] - 2025-01-XX

#### Added
- Initial release of the command-line profit calculator
- Basic profit calculation functionality
- Support for revenue, expenses, and tax rate inputs
- EBT (Earnings Before Tax) calculation
- Profit ratio calculation
- Output in both cents and euros

## 🐛 Issues

If you encounter any issues or have suggestions for improvements, please [open an issue](https://github.com/vahiiiid/profit-calculator/issues) on GitHub.

When reporting GUI-related issues, please include:
- Your operating system and version
- Go version
- Steps to reproduce the issue
- Screenshots if applicable

## 👨‍💻 Author

**vahiiiid** - [GitHub Profile](https://github.com/vahiiiid)

---

⭐ If you find this project helpful, please consider giving it a star on GitHub!

### Main Interface
The clean and intuitive interface makes profit calculations simple and efficient:

*The application features three input fields for revenue, expenses, and tax rate, with Calculate and Reset buttons for easy operation.* 