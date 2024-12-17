import React from 'react';
import './ListContainer.css';

const ListContainer: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <div className="list-container">
      {children}
    </div>
  );
};

export default ListContainer;