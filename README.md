
# LlamaShell

<div align="center">
	<img width=250 src="assets/logo.png">
</div>

----

**LlamaShell** A simple Retrieval-Augmented Generation (RAG) system leveraging the Llama 3.1 model to enhance Linux CLI usability. Implemented using Go and Ollama, LlamaShell provides context-aware command suggestions and automation to streamline developer workflows.

## ⚡ Features

-   **Context-Aware Suggestions**: Get real-time, intelligent suggestions based on your CLI commands and context.
-   **Powered by Ollama**: Utilizes Ollama models for optimized performance and scalability tailored specifically for Llama-based models.


## 🎉 Getting Started

### Prerequisites

-   **Go**: Make sure you have Go installed. You can download it [here](https://golang.org/dl/).
-   **Ollama**: Set up Ollama models on your machine or server. Refer to the [Ollama](https://ollama.com/) documentation for installation instructions.
-   **LangChain Go**: Install the LangChain Go library. More information is available in the [LangChain Go documentation](https://github.com/langchain-ai/langchain-go).

### Installation

1.  **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/LlamaShell.git 
    cd LlamaShell
    ```
2. **Install Dependencies**: Ensure that all necessary Go modules and dependencies are installed:
	```bash
	go mod tidy
3.   **Set Up Ollama Models**: Follow the instructions in the Ollama documentation to set up the required models for LlamaShell.
    
### License

This project is licensed under the MIT License. See the LICENSE file for details.

### 🔗 Acknowledgments

-   **LangChain Go**
-   **Ollama**
-   **Cobra CLI**