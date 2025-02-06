package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

// Reciever
type TextEditor struct {
	text string
}

// Adds new text to the editor
func (e *TextEditor) AddText(newText string) {
	e.text += newText

}

// Removes the last added text
func (e *TextEditor) RemoveText(length int) {
	if len(e.text) >= length {
		e.text = e.text[:len(e.text)-length]
	}
}

// Get Text
func (e *TextEditor) GetText() string {
	return e.text
}

// Concrete Command
type AddTextCommand struct {
	editor *TextEditor
	text   string
}

func (c *AddTextCommand) Execute() {
	c.editor.AddText(c.text)

}
func (c *AddTextCommand) Undo() {
	c.editor.RemoveText(len(c.text))
}

// Invoker
type CommandManager struct {
	history   []Command
	redoStack []Command
}

// ExecuteCommand runs a command and stores it
func (m *CommandManager) ExecuteCommand(cmd Command) {
	cmd.Execute()
	m.history = append(m.history, cmd)
	m.redoStack = nil
}

// Undo reverts the last command
func (m *CommandManager) Undo() {
	if len(m.history) == 0 {
		fmt.Println("Nothing to Undo")
		return
	}
	lastCommand := m.history[len(m.history)-1]
	m.history = m.history[:len(m.history)-1]
	lastCommand.Undo()
	m.redoStack = append(m.redoStack, lastCommand)
}

// Redo reapplies the last undone command
func (m *CommandManager) Redo() {
	if len(m.redoStack) == 0 {
		fmt.Println("Nothing to Redo")
		return
	}

	lastCommand := m.redoStack[len(m.redoStack)-1]
	m.redoStack = m.redoStack[:len(m.redoStack)-1]
	lastCommand.Execute()
	m.history = append(m.history, lastCommand)
}

func main() {

	editor := new(TextEditor)
	manager := new(CommandManager)

	// Excute typing commands
	cmd1 := &AddTextCommand{editor: editor, text: "Hello "}
	cmd2 := &AddTextCommand{editor: editor, text: "World! "}
	manager.ExecuteCommand(cmd1)
	manager.ExecuteCommand(cmd2)

	fmt.Println("Current Text:", editor.GetText())

	manager.Undo()
	fmt.Println("After Undo:", editor.GetText()) // "Hello"

	manager.Redo()
	fmt.Println("After Redo:", editor.GetText()) // "Hello World!"

}
