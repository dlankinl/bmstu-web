import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';
import Modal from '../Modal/Modal';
import SignupComponent from '../Signup/Signup';
import SigninComponent from '../Signin/Signin';

const Navbar = () => {
  // const [isModalActive, setModalActive] = useState(false);
    
  // const handleModalOpen = () => {
  //   setModalActive(true);
  // };
  // const handleModalClose = () => {
  //   setModalActive(false);
  // };

  const [isModalActive, setModalActive] = useState(false);
  const [modalContent, setModalContent] = useState<React.ReactNode>(null); // State to track modal content

  const handleModalOpen = (content: React.ReactNode) => {
    setModalContent(content); // Set the content based on button clicked
    setModalActive(true);
  };

  const handleModalClose = () => {
    setModalActive(false);
    setModalContent(null); // Reset content when closing
  };

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        {/* <img src="path_to_your_logo.png" alt="Logo" /> */}
        <button className="nav-button">VentureContact</button>
      </div>
      <div className="navbar-buttons">
        <button className="nav-button">Найти партнера</button>
        <Link to="/activity-fields">
          <button className="nav-button">Сферы деятельности</button>
        </Link>
        <div className="auth-buttons">
          <button className="auth-button" onClick={() => { handleModalOpen(<SignupComponent/>) }}>Зарегистрироваться</button>
          <button className="auth-button" onClick={() => { handleModalOpen(<SigninComponent/>) }}>Войти</button>
        </div>
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