import React from 'react';
import './GreenButton.css';

const GreenButton = ({ text, icon }) => {
    return (
        <button type="submit" className="icon-button">
            {icon && <span className="icon">{icon}</span>}
            <span>{text}</span>
        </button>
    );
};

// Export the component
export default GreenButton;