import React from 'react';
import { Entrepreneur } from './types';
import ListItem from '../List/Item';
import { Link } from 'react-router-dom';
import "./Entrepreneur.css";

const EntrepreneurItem: React.FC<{ entrepreneur: Entrepreneur }> = ({ entrepreneur }) => {
  const formattedRating = entrepreneur.Rating.toFixed(1);

  return (
    <ListItem
      item={entrepreneur}
      renderContent={(item) => (
        <>
          <div className="item-details">
            <div className="item-info">
              <div className="item-rating">
                <Link to={`/entrepreneurs/${item.ID}`}>
                  <h3>{item.Name}</h3>
                </Link>
                <span className="rating-circle">{formattedRating}</span> 
              </div>
              <p>{item.Gender} | {item.Birthday}</p>
            </div>
            <h2 className="item-description">{item.City}</h2>
          </div>
        </>
      )}
    />
  );
};

export default EntrepreneurItem;