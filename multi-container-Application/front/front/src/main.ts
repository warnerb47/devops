import './style.css';

type Todo = {
  id: number;
  text: string;
  completed: boolean;
};

class TodoApp {
  private todos: Todo[] = [];
  private todoInput: HTMLInputElement;
  private todoList: HTMLUListElement;
  private todoCount: HTMLElement;
  private filterButtons: NodeListOf<HTMLButtonElement>;
  private currentFilter: 'all' | 'active' | 'completed' = 'all';

  constructor() {
    this.todos = this.loadTodos();
    this.initElements();
    this.setupEventListeners();
    this.render();
  }

  private initElements(): void {
    const app = document.querySelector<HTMLDivElement>('#app')!;
    app.innerHTML = `
      <div class="todo-container">
        <h1>Todo App</h1>
        <div class="todo-input-container">
          <input 
            type="text" 
            id="todo-input" 
            placeholder="What needs to be done?"
            class="todo-input"
          />
          <button id="add-todo" class="add-button">Add</button>
        </div>
        <div class="filters">
          <button data-filter="all" class="filter-btn active">All</button>
          <button data-filter="active" class="filter-btn">Active</button>
          <button data-filter="completed" class="filter-btn">Completed</button>
        </div>
        <ul id="todo-list" class="todo-list"></ul>
        <div class="todo-footer">
          <span id="todo-count">0 items left</span>
          <button id="clear-completed" class="clear-btn">Clear Completed</button>
        </div>
      </div>
    `;

    this.todoInput = document.getElementById('todo-input') as HTMLInputElement;
    this.todoList = document.getElementById('todo-list') as HTMLUListElement;
    this.todoCount = document.getElementById('todo-count') as HTMLElement;
    this.filterButtons = document.querySelectorAll('.filter-btn');
  }

  private setupEventListeners(): void {
    // Add todo on button click or Enter key
    document.getElementById('add-todo')?.addEventListener('click', () => this.addTodo());
    this.todoInput?.addEventListener('keypress', (e) => {
      if (e.key === 'Enter') this.addTodo();
    });

    // Filter todos
    this.filterButtons.forEach(button => {
      button.addEventListener('click', () => this.setFilter(button.dataset.filter as any));
    });

    // Clear completed todos
    document.getElementById('clear-completed')?.addEventListener('click', () => this.clearCompleted());
  }

  private addTodo(): void {
    const text = this.todoInput.value.trim();
    if (text) {
      const newTodo: Todo = {
        id: Date.now(),
        text,
        completed: false
      };
      this.todos.push(newTodo);
      this.saveTodos();
      this.todoInput.value = '';
      this.render();
    }
  }

  private toggleTodo(id: number): void {
    const todo = this.todos.find(t => t.id === id);
    if (todo) {
      todo.completed = !todo.completed;
      this.saveTodos();
      this.render();
    }
  }

  private deleteTodo(id: number): void {
    this.todos = this.todos.filter(todo => todo.id !== id);
    this.saveTodos();
    this.render();
  }

  private clearCompleted(): void {
    this.todos = this.todos.filter(todo => !todo.completed);
    this.saveTodos();
    this.render();
  }

  private setFilter(filter: 'all' | 'active' | 'completed'): void {
    this.currentFilter = filter;
    this.filterButtons.forEach(btn => {
      if (btn.dataset.filter === filter) {
        btn.classList.add('active');
      } else {
        btn.classList.remove('active');
      }
    });
    this.render();
  }

  private getFilteredTodos(): Todo[] {
    switch (this.currentFilter) {
      case 'active':
        return this.todos.filter(todo => !todo.completed);
      case 'completed':
        return this.todos.filter(todo => todo.completed);
      default:
        return [...this.todos];
    }
  }

  private saveTodos(): void {
    localStorage.setItem('todos', JSON.stringify(this.todos));
  }

  private loadTodos(): Todo[] {
    const saved = localStorage.getItem('todos');
    return saved ? JSON.parse(saved) : [];
  }

  private render(): void {
    const filteredTodos = this.getFilteredTodos();
    
    this.todoList.innerHTML = filteredTodos.map(todo => `
      <li class="todo-item ${todo.completed ? 'completed' : ''}" data-id="${todo.id}">
        <input 
          type="checkbox" 
          ${todo.completed ? 'checked' : ''} 
          class="todo-checkbox"
        />
        <span class="todo-text">${todo.text}</span>
        <button class="delete-btn">×</button>
      </li>
    `).join('');

    // Update todo count
    const activeCount = this.todos.filter(todo => !todo.completed).length;
    this.todoCount.textContent = `${activeCount} ${activeCount === 1 ? 'item' : 'items'} left`;

    // Add event listeners to todo items
    document.querySelectorAll('.todo-item').forEach(item => {
      const id = parseInt(item.getAttribute('data-id') || '0');
      const checkbox = item.querySelector('.todo-checkbox') as HTMLInputElement;
      const deleteBtn = item.querySelector('.delete-btn') as HTMLButtonElement;
      
      checkbox.addEventListener('change', () => this.toggleTodo(id));
      deleteBtn.addEventListener('click', () => this.deleteTodo(id));
    });
  }
}

// Initialize the app
new TodoApp();
