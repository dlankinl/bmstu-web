import React from 'react';
import './FormContainer.css';

const FormContainer: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <div className="form-container">
      {children}
    </div>
  );
};

export default FormContainer;