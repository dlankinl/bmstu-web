import React from 'react';
import Modal from 'react-modal';

interface CustomModalProps {
  isOpen: boolean;
  onRequestClose: () => void;
  children: React.ReactNode;
}

const CustomModal: React.FC<CustomModalProps> = ({ isOpen, onRequestClose, children }) => {
  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onRequestClose}
      style={{
        overlay: {
          backgroundColor: 'rgba(0, 0, 0, 0.75)',
        },
        content: {
          color: 'lightsteelblue',
        },
      }}
      contentLabel="Activity Field Modal"
    >
      <h2>Add Activity Field</h2>
      <button onClick={onRequestClose}>Close</button>
      <div>{children}</div>
    </Modal>
  );
};

export default CustomModal;