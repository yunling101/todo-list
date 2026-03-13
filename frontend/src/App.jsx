import { useState, useEffect } from 'react';
import TodoList from './components/TodoList';
import TodoForm from './components/TodoForm';
import { fetchTodos, createTodo, updateTodo, deleteTodo, toggleTodo } from './api/todo';
import './App.css';

function App() {
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    loadTodos();
  }, []);

  const loadTodos = async () => {
    try {
      setLoading(true);
      const data = await fetchTodos();
      setTodos(data);
      setError(null);
    } catch (err) {
      setError('无法加载待办事项，请检查后端服务是否运行');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = async (title) => {
    try {
      const newTodo = await createTodo(title);
      setTodos([...todos, newTodo]);
    } catch (err) {
      setError('添加失败');
      console.error(err);
    }
  };

  const handleEdit = async (id, title, completed) => {
    try {
      const updatedTodo = await updateTodo(id, title, completed);
      setTodos(todos.map((t) => (t.id === id ? updatedTodo : t)));
    } catch (err) {
      setError('更新失败');
      console.error(err);
    }
  };

  const handleDelete = async (id) => {
    try {
      await deleteTodo(id);
      setTodos(todos.filter((t) => t.id !== id));
    } catch (err) {
      setError('删除失败');
      console.error(err);
    }
  };

  const handleToggle = async (id) => {
    try {
      const updatedTodo = await toggleTodo(id);
      setTodos(todos.map((t) => (t.id === id ? updatedTodo : t)));
    } catch (err) {
      setError('操作失败');
      console.error(err);
    }
  };

  return (
    <div className="app">
      <div className="container">
        <header className="header">
          <h1>📝 待办事项</h1>
        </header>

        <TodoForm onAdd={handleAdd} />

        {error && (
          <div className="error-message">
            <span>{error}</span>
            <button onClick={() => setError(null)}>×</button>
          </div>
        )}

        {loading ? (
          <div className="loading">
            <div className="spinner"></div>
            <p>加载中...</p>
          </div>
        ) : (
          <TodoList
            todos={todos}
            onToggle={handleToggle}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        )}
      </div>
    </div>
  );
}

export default App;
