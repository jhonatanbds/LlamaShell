
# LlamaShell (WIP)

<div align="center">
	<img width=250 src="assets/logo.png">
</div>

----

🦙 **LlamaShell** is a simple Retrieval-Augmented Generation (RAG) system leveraging the Llama 3.1 model to enhance Linux CLI usability. Implemented using Go and Ollama, LlamaShell provides context-aware command suggestions and automation to streamline developer workflows.

## ⚡ Features

-   **Context-Aware Suggestions**: Get real-time, intelligent suggestions based on your CLI commands and context.
-   **Powered by Ollama**: Utilizes Ollama models for optimized performance and scalability tailored specifically for Llama-based models.


## 🎉 Getting Started

### Prerequisites

-   **Go**: Make sure you have Go installed. You can download it [here](https://golang.org/dl/).
-   **Ollama**: Set up Ollama models on your machine or server. Refer to the [Ollama](https://ollama.com/) documentation for installation instructions. Typically:
    ```bash
    curl -fsSL https://ollama.com/install.sh | sh
    ```
-   **LangChain Go**: Install the LangChain Go library. More information is available in the [LangChain Go documentation](https://github.com/langchain-ai/langchain-go).

### Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/LlamaShell.git 
    cd LlamaShell
    ```
2. **Install Dependencies**: Ensure that all necessary Go modules and dependencies are installed:
	```bash
	go mod tidy
3. **Build and install**:
    ```bash
    go build -o llamash
    go install
    ```


## Usage

### Chat

Activate LlamaShell CLI with `llamash chat`, it will open the chat prompt where LlamaShell will suggest commands for your tasks. For instance, see the example below:

```bash
Chat with 🦙 LlamaShell. Type 'exit' to end the conversation.
>>> I need to create a text file with 36 random numbers
dd if=/dev/urandom bs=1 count=36 2>/dev/null | tr -dc '0-9' > random_numbers.txt
>>> upload this file to the gcp bucket gs://storage/personal
gsutil cp random_numbers.txt gs://storage/personal
>>> I wanna move the file to a new bucket: gs://storage/personal/random
gsutil mv gs://storage/personal/random_numbers.txt gs://storage/personal/random/
>>> exit
Ending chat.
```

### Run

Activate LlamaShell CLI with `llamash run <your desired task>`, LlamaShell will show you the suggested command to perform the desired task, which you can run or regenerate the suggested command for another option. Use the arrow keys to navigate up and down, `Enter` to choose or press `q` to quit. For instance, see the example below:

```bash
llamash run make a hello world python file 
> echo "print('Hello, World!')" > hello.py
  Regenerate
llamash run show the content of the file hello.py
> cat hello.py
  Regenerate
print('Hello, World!')
```

### License

This project is licensed under the MIT License. See the LICENSE file for details.

### 🔗 Acknowledgments

-   [**LangChain Go**](https://github.com/tmc/langchaingo)
-   [**Ollama**](https://ollama.com/)
-   [**Cobra CLI**](https://github.com/spf13/cobra)

## TODO

 - [X] chat cmd
 - [X] loading CLI animation
 - [X] LLM initial context
 - [X] run cmd
 - [ ] set configs: ollama model
 - [ ] create config file
 - [x] build + install
 - [ ] run option to chat cmd
 - [ ] RAG
 - [ ] use command history context
 - [ ] use directory tree context
 - [ ] tests :cry: