# Simulação de Mercado de Ações

Este projeto é um sistema de simulação de mercado de ações desenvolvido em **Go**, com suporte a **concorrência**, **banco de dados** e **testes automatizados**.

## 📌 Funcionalidades

### 1️⃣ Usuários
- Criar novos usuários com saldo inicial.
- Buscar informações de um usuário específico.

### 2️⃣ Ações
- Registrar ações disponíveis para compra e venda.
- Consultar ações disponíveis no mercado.

### 3️⃣ Transações
- Permitir que usuários comprem e vendam ações.
- Atualizar automaticamente o saldo dos usuários e a quantidade de ações possuídas.

## 🔧 Tecnologias Utilizadas
- **Go**
- **Banco de Dados (SQLite em memória para testes, PostgreSQL/MySQL para produção)**
- **Testes unitários com `testing`**
- **Concorrência para simular múltiplos usuários operando ao mesmo tempo**

## 🚀 Objetivo
Este projeto tem como foco aprimorar as habilidades em **Go**, explorando a implementação de **banco de dados**, **concorrência** e **boas práticas de testes**.

---

📢 **Como Rodar?**
1. Clone este repositório.
2. Configure o banco de dados no arquivo `.env`.
3. Execute as migrations.
4. Rode a aplicação e os testes!

Se precisar de mais informações, sinta-se à vontade para contribuir ou sugerir melhorias! 🚀