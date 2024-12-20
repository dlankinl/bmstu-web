import React from 'react';
import './GreenButton.css';

interface GreenButtonProps {
  text: string;
  icon?: string; // Optional icon prop
  onClick?: () => void; // Add onClick prop
}

// const GreenButton = ({ text, icon }) => {
  // return (
  //   <button type="submit" className="icon-button">
  //     {icon && <span className="icon">{icon}</span>}
  //     <span>{text}</span>
  //   </button>
  // );
// };

const GreenButton: React.FC<GreenButtonProps> = ({ text, icon, onClick }) => {
  return (
    <button type="submit" className="icon-button" onClick={onClick}>
      {icon && <span className="icon">{icon}</span>}
      <span>{text}</span>
    </button>
  );
};


// Export the component
export default GreenButton;