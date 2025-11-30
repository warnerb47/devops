const API_BASE_URL = 'http://localhost:8080/v1/todos';

export interface Todo {
  id: string;
  label: string;
  checked: boolean;
  date?: string;
}

export const todoApi = {
  // Get all todos
  async getAll(): Promise<Todo[]> {
    const response = await fetch(API_BASE_URL);
    if (!response.ok) {
      throw new Error('Failed to fetch todos');
    }
    return response.json();
  },

  // Get a single todo by ID
  async getById(id: string): Promise<Todo> {
    const response = await fetch(`${API_BASE_URL}/${id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch todo with id ${id}`);
    }
    return response.json();
  },

  // Create a new todo
  async create(todo: Omit<Todo, 'id' | 'date'>): Promise<Todo> {
    const response = await fetch(API_BASE_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(todo),
    });
    if (!response.ok) {
      throw new Error('Failed to create todo');
    }
    return response.json();
  },

  // Update an existing todo
  async update(id: string, todo: Partial<Todo>): Promise<Todo> {
    const response = await fetch(`${API_BASE_URL}/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(todo),
    });
    if (!response.ok) {
      throw new Error(`Failed to update todo with id ${id}`);
    }
    return response.json();
  },

  // Delete a todo
  async delete(id: string): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error(`Failed to delete todo with id ${id}`);
    }
  }
};
