# DataForgeOS

Sistema **open source** educacional que simula um ambiente empresarial integrado para aprendizado prático em programação, engenharia de dados e desenvolvimento backend.

---

## 🎯 Objetivos

- Fornecer um ambiente realista de prática para estudantes e profissionais.  
- Permitir interação manual via interface gráfica, automação de dados e consumo de APIs.  
- Ensinar boas práticas de arquitetura de software, design de banco de dados e integração entre sistemas.  

---

## 🛠️ Arquitetura

O DataForgeOS é **modular**, composto por três subsistemas principais:

| Módulo            | Descrição                                |
|-------------------|------------------------------------------|
| **Produção**      | Registra dados de produção por turno.    |
| **Produtos**  | Gerencia cadastro e atributos de Produtos. |
| **Vendas**        | Controla transações de vendas.           |

Cada módulo contém **três componentes**:

1. **Interface Gráfica (GUI)** – Go + Fyne (desktop).  
2. **API REST** – Go (Gin ou Fiber) para CRUD e relatórios.  
3. **Bot de Dados** – Python (Faker, Schedule, Requests) para gerar dados fictícios e popular o PostgreSQL.
4. **Banco de dados** - PostgreSQL

---

## 🔗 Integração

- **Banco de Dados:** PostgreSQL (esquema multi‑módulo).  
- **Comunicação:** API REST entre GUI e banco.  

---

## 🚀 Pré‑requisitos

| Ferramenta | Versão mínima |
|------------|---------------|
| Go         | 1.20          |
| Python     | 3.9           |
| PostgreSQL | 12            |
| make       | (opcional)    |

---

```bash
# Clone o repositório
git clone https://github.com/oBrenn0w/DataForgeOS.git
cd DataForgeOS

# Configure o banco de dados
psql -U <usuario> -d <database> -f migrations/init.sql

# Instale dependências Go
go mod download

# Instale dependências do Bot Python
cd bot
pip install -r requirements.txt

# Inicie a API REST
cd ../api
go run main.go

# Inicie a GUI desktop
cd ../gui
go run main.go

```
---
```text
DataForgeOS/
├── api/           # Código da API REST em Go
├── gui/           # Interface Fyne em Go
├── bot/           # Bot de dados em Python
├── migrations/    # Scripts SQL para criação de esquema
├── docs/          # Documentação adicional
└── README.md      # Visão geral do projeto

