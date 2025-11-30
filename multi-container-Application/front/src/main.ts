import './style.css';
import { todoApi, type Todo } from './api';

class TodoApp {
  private todos: Todo[] = [];
  private todoInput!: HTMLInputElement;
  private todoList!: HTMLUListElement;
  private todoCount!: HTMLElement;
  private filterButtons!: NodeListOf<HTMLButtonElement>;
  private currentFilter: 'all' | 'active' | 'completed' = 'all';
  private isLoading: boolean = false;

  constructor() {
    this.initElements();
    this.setupEventListeners();
    this.loadEnvConfig().then(
      () => this.loadTodos()
    );
  }

  private async loadEnvConfig() {
    const response = await fetch('/env.json');
    if (!response.ok) {
      throw new Error('Failed to fetch env');
    }
    const env = await response.json();
    console.log({env});
    todoApi.apiBaseUrl = env.API_BASE_URL;
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

  private async addTodo(): Promise<void> {
    const label = this.todoInput.value.trim();
    if (label) {
      try {
        this.isLoading = true;
        const newTodo = await todoApi.create({
          label,
          checked: false
        });
        this.todos.push(newTodo);
        this.todoInput.value = '';
        this.isLoading = false;
        this.render();
      } catch (error) {
        console.error('Failed to add todo:', error);
        alert('Failed to add todo. Please try again.');
      } finally {
        this.isLoading = false;
      }
    }
  }

  private async toggleTodo(id: string): Promise<void> {
    const todo = this.todos.find(t => t.id === id);
    if (todo) {
      try {
        this.isLoading = true;
        const updatedTodo = await todoApi.update(id, {
          checked: !todo.checked
        });
        // Update the local todo with the server response
        Object.assign(todo, updatedTodo);
        this.isLoading = false;
        this.render();
      } catch (error) {
        console.error('Failed to update todo:', error);
        alert('Failed to update todo. Please try again.');
      } finally {
        this.isLoading = false;
      }
    }
  }

  private async deleteTodo(id: string): Promise<void> {
    try {
      this.isLoading = true;
      await todoApi.delete(id);
      this.todos = this.todos.filter(todo => todo.id !== id);
      this.isLoading = false;
      this.render();
    } catch (error) {
      console.error('Failed to delete todo:', error);
      alert('Failed to delete todo. Please try again.');
    } finally {
      this.isLoading = false;
    }
  }

  private async clearCompleted(): Promise<void> {
    try {
      this.isLoading = true;
      // Delete all completed todos
      const deletePromises = this.todos
        .filter(todo => todo.checked)
        .map(todo => todoApi.delete(todo.id));
      
      await Promise.all(deletePromises);
      
      // Update local state
      this.todos = this.todos.filter(todo => !todo.checked);
      this.isLoading = false;
      this.render();
    } catch (error) {
      console.error('Failed to clear completed todos:', error);
      alert('Failed to clear completed todos. Please try again.');
    } finally {
      this.isLoading = false;
    }
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
        return this.todos.filter(todo => !todo.checked);
      case 'completed':
        return this.todos.filter(todo => todo.checked);
      default:
        return [...this.todos];
    }
  }

  private async loadTodos(): Promise<void> {
    try {
      this.isLoading = true;
      this.todos = await todoApi.getAll();
      this.isLoading = false;
      this.render();
    } catch (error) {
      console.error('Failed to load todos:', error);
      alert('Failed to load todos. Please try again later.');
    } finally {
      this.isLoading = false;
    }
  }

  private render(): void {
    const filteredTodos = this.getFilteredTodos();
    
    this.todoList.innerHTML = this.isLoading 
      ? '<div class="loading">Loading...</div>'
      : filteredTodos.map(todo => `
        <li class="todo-item ${todo.checked ? 'completed' : ''}" data-id="${todo.id}">
          <input 
            type="checkbox" 
            ${todo.checked ? 'checked' : ''} 
            class="todo-checkbox"
          />
          <span class="todo-text">${todo.label}</span>
          <button class="delete-btn">×</button>
        </li>
      `).join('');

    // Update todo count
    const activeCount = this.todos.filter(todo => !todo.checked).length;
    this.todoCount.textContent = `${activeCount} ${activeCount === 1 ? 'item' : 'items'} left`;

    // Add event listeners to todo items
    document.querySelectorAll('.todo-item').forEach(item => {
      const id = item.getAttribute('data-id') || '';
      const checkbox = item.querySelector('.todo-checkbox') as HTMLInputElement;
      const deleteBtn = item.querySelector('.delete-btn') as HTMLButtonElement;
      
      checkbox.addEventListener('change', () => this.toggleTodo(id));
      deleteBtn.addEventListener('click', () => this.deleteTodo(id));
    });
  }
}

// Initialize the app
new TodoApp();
