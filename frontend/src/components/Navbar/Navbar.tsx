import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';
import Modal from '../Modal/Modal';
import SignupComponent from '../Signup/Signup';
import SigninComponent from '../Signin/Signin';

const Navbar = () => {
  const [isModalActive, setModalActive] = useState(false);
  const [modalContent, setModalContent] = useState<React.ReactNode>(null); 

  const handleModalOpen = (content: React.ReactNode) => {
    setModalContent(content);
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
    setModalContent(null);
    window.location.reload();
  };

  const isAuthenticated = !!localStorage.getItem('authToken');
  const handleLogout = () => {
    localStorage.removeItem('authToken');
    window.location.reload();
  };

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <button className="nav-button">VentureContact</button>
      </div>
      <div className="navbar-buttons">
        <Link to="/entrepreneurs">
          <button className="nav-button">Найти партнера</button>
        </Link>
        <Link to="/activity-fields">
          <button className="nav-button">Сферы деятельности</button>
        </Link>
        {!isAuthenticated ? (
          <div className="auth-buttons">
            <button className="auth-button" onClick={() => { handleModalOpen(<SignupComponent onSuccess={handleModalClose} />) }}>Зарегистрироваться</button>
            <button className="auth-button" onClick={() => { handleModalOpen(<SigninComponent onSuccess={handleModalClose} />) }}>Войти</button>
          </div>
        ) : (
          <div className="auth-buttons">
            <button className="auth-button" onClick={handleLogout}>Выйти</button>
          </div>
        )}
      </div>
      <div>
        {isModalActive && (
          <Modal onClose={handleModalClose}>
            {modalContent}
          </Modal>
        )}
      </div>
    </nav>
  );
};

export default Navbar;