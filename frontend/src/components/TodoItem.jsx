import { useState } from 'react';

const TodoItem = ({ todo, onToggle, onEdit, onDelete }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editTitle, setEditTitle] = useState(todo.title);

  const handleSave = () => {
    if (editTitle.trim()) {
      onEdit(todo.id, editTitle.trim(), todo.completed);
      setIsEditing(false);
    }
  };

  const handleCancel = () => {
    setEditTitle(todo.title);
    setIsEditing(false);
  };

  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      handleSave();
    } else if (e.key === 'Escape') {
      handleCancel();
    }
  };

  return (
    <li className={`todo-item ${todo.completed ? 'completed' : ''}`}>
      {isEditing ? (
        <div className="edit-form">
          <input
            type="text"
            value={editTitle}
            onChange={(e) => setEditTitle(e.target.value)}
            onKeyDown={handleKeyDown}
            autoFocus
            maxLength={100}
          />
          <div className="edit-actions">
            <button className="save-btn" onClick={handleSave}>保存</button>
            <button className="cancel-btn" onClick={handleCancel}>取消</button>
          </div>
        </div>
      ) : (
        <>
          <div className="todo-content" onClick={() => onToggle(todo.id)}>
            <input
              type="checkbox"
              checked={todo.completed}
              onChange={() => onToggle(todo.id)}
            />
            <span className="todo-title">{todo.title}</span>
          </div>
          <div className="todo-actions">
            <button
              className="edit-btn"
              onClick={() => setIsEditing(true)}
              title="编辑"
            >
              ✏️
            </button>
            <button
              className="delete-btn"
              onClick={() => onDelete(todo.id)}
              title="删除"
            >
              🗑️
            </button>
          </div>
        </>
      )}
    </li>
  );
};

export default TodoItem;
