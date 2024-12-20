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
      appElement={document.getElementById('root')}
      style={{
        overlay: {
          backgroundColor: 'rgba(0, 0, 0, 0.75)',
        },
        content: {
          display: 'flex', 
          justifyContent: 'center',
          alignItems: 'center', 
          padding: 0, 
          border: 'none',
          backgroundColor: 'transparent', 
          marginTop: '100px'
        },
      }}
    >
      <div style={{ width: '500px', height: '100%' }}>{children}</div>
    </Modal>
  );
};

export default CustomModal;