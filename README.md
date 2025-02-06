# 🚀 Design Patterns in Golang

This repository explores various **Design Patterns** in Golang, categorized into **Creational, Structural, and Behavioral** patterns. These patterns help in **object creation, structuring code efficiently, and improving object interaction**.

---

## 📌 1. Creational Patterns (Object Creation)
Creational patterns focus on **how objects are created**, ensuring that the instantiation process is efficient, scalable, and maintainable.

### 🏆 1.1 Singleton Pattern
✔ Ensures **only one instance** of a struct exists.  
✔ **Thread-safe**, preventing race conditions in concurrent environments.  
✔ Example use case: **Database connections, configuration managers.**  

### 🏭 1.2 Factory Pattern
✔ Abstracts the object creation logic, improving **maintainability and extensibility**.  
✔ Helps in creating different object types based on input parameters.  
✔ Example use case: **Loggers, payment gateways, and parsers.**

---

## 📌 2. Structural Patterns (Object Composition)
Structural patterns focus on **how objects and classes are composed**, ensuring better **code organization, flexibility, and reuse**.

### 🔌 2.1 Adapter Pattern
✔ **Bridges incompatible interfaces**, making integration seamless.  
✔ Useful when **modernizing old APIs** for new systems.  
✔ Example use case: **Connecting legacy payment APIs with new transaction systems.**  

### 🎭 2.2 Decorator Pattern
✔ Dynamically **adds features to objects** without modifying the original struct.  
✔ Keeps the core logic clean and reusable.  
✔ Example use case: **Middleware in HTTP servers (logging, authentication, compression).**  

### 🛡 2.3 Proxy Pattern
✔ Acts as a **wrapper** for another object to provide **security, caching, or lazy loading**.  
✔ Reduces resource usage by controlling access to expensive operations.  
✔ Example use case: **Lazy database connections, API rate limiting, and security layers.**  

---

## 📌 3. Behavioral Patterns (Object Interaction)
Behavioral patterns focus on **how objects communicate and interact**, making the system **more flexible and loosely coupled**.

### 🔄 3.1 Strategy Pattern
✔ **Dynamically switches between different algorithms or strategies** at runtime.  
✔ Promotes the **Open/Closed Principle** (adding new strategies without modifying existing code).  
✔ Example use case: **Sorting algorithms, payment processing (Card, UPI, PayPal).**  

### 📢 3.2 Observer Pattern
✔ Implements **event-driven systems**, where multiple objects react to changes in one object.  
✔ Useful when a **one-to-many dependency** exists between objects.  
✔ Example use case: **Stock market updates, messaging queues, notification systems.**  

### 🎮 3.3 Command Pattern
✔ Encapsulates **requests as objects**, allowing easy execution, undo, and redo operations.  
✔ Helps in implementing **action queues, macros, and transaction rollbacks**.  
✔ Example use case: **Undo/redo in text editors, remote controls, task scheduling.**  

---

## 🔗 **Want to Learn More?**
- Read **Golang Design Patterns Documentation** 📖 [here](https://golang.org/doc/)
- Explore real-world **design pattern implementations** in Go 🛠

---

## 🤝 **Contributions & Feedback**
Feel free to contribute by improving implementations or adding more patterns!  
Found an issue? Open a ticket or reach out. 🚀
