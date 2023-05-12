# RegexAI

RegexAI is a command-line tool built with Go that leverages the OpenAI GPT-3 model to generate regex patterns. It provides an intuitive interface to generate regex patterns right from your terminal.

## Requirements
* **Go**: Ensure that you have Go 1.18 or later installed on your system. You can check your Go installation by running `go version` in your terminal. If you don't have Go installed or have an older version, you can download the latest version from [the official Go website](https://golang.org/dl/).
* **OpenAI API Key**: You can obtain an API key from [OpenAI website](https://platform.openai.com/).

## Installation
To install RegexAI, follow the steps below:

1. Clone the RegexAI repository to your local machine:

    ```bash
    git clone https://github.com/bugfloyd/regexai.git
    ```

2. Navigate to the cloned directory:

    ```bash
    cd regexai
    ```

3. Copy the example .env file to `.env`:
    ```bash
   cp .env.example .env
    ```

4. Update the content of the `.env` file with your OpenAI API key. See the Configuration section for more details.


5. Build the project:

    ```bash
    go build
    ```

   This will generate a binary named `regexai`.

6. Add execution permissions to the binary:

    ```bash
    chmod +x regexai
    ```

## Usage

RegexAI is very simple to use. Here's the basic usage:

```bash
./regexai "email address"
```

This will generate a regex pattern for validating an email address.

If you'd like a short explanation about the generated regex, you can use the `-e` or `--explain` flag:

```bash
./regexai -e "email address"
```

Or:

```bash
./regexai --explain "email address"
```

This will generate a regex pattern along with a short explanation.

## Configuration

RegexAI uses the OpenAI API, and it requires an API key to function. The API key should be set in a `.env` file in the root of the project. Here's a sample `.env` file:

```env
OPENAI_API_KEY=yourapikey
```

Replace `yourapikey` with your actual OpenAI API key.

## Contributing

We welcome contributions from the community. If you'd like to contribute, please fork the repository and make your changes, then create a pull request against the main branch.

## License

RegexAI is licensed under the GNU General Public License v2.0. See `LICENSE` for more details.

## Disclaimer

This tool uses the OpenAI GPT-3 model, which is a powerful language model. However, it's not perfect, and it may not always generate the correct regex pattern. Always test the generated regex before using it in production.