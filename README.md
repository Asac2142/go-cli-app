<h1 align="center"><strong>Task Tracker CLI</strong></h1>

## Description
Task tracker is a project used to track and manage your tasks written in Go.
https://roadmap.sh/projects/task-tracker

## Usage 🛠
- `git clone https://github.com/Asac2142/tracker-cli-go.git`
- `cd tracker-cli-go`
- `go build -o task-cli`

## Add task ➕
```
task-cli add <task-description>
```

## Update task 🔄
```
task-cli update <task-id> <task-new-description>
```

## Delete task 🗑️
```
task-cli delete <task-id>
```

## Update task status to in-progress 🔧
```
task-cli mark-in-progress <task-id>
```
## Update task status to done ✅
```
task-cli mark-done <task-id>`
```

## Listing tasks 👀
```
task-cli list             //list all tasks
task-cli list done        //list tasks with status done
task-cli list todo        //list tasks with status todo
task-cli list in-progress //list tasks with status in-progress
```
