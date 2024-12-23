import React from 'react';
import { Entrepreneur } from './types';
import ListItem from '../List/Item';
import { Link } from 'react-router-dom';
import "./Entrepreneur.css";

const EntrepreneurItem: React.FC<{ entrepreneur: Entrepreneur }> = ({ entrepreneur }) => {
  const formattedRating = 0.0;
  if (entrepreneur.rating) {
    const formattedRating = entrepreneur.rating.toFixed(1);
  }

  return (
    <ListItem
      item={entrepreneur}
      renderContent={(item) => (
        <>
          <div className="item-details">
            <div className="item-info">
              <div className="item-rating">
                <Link to={`/entrepreneurs/${item.id}`}>
                  <h3>{item.fullName}</h3>
                </Link>
                <span className="rating-circle">{formattedRating}</span> 
              </div>
              <p>
              {item.gender === 'm' ? 'мужской ' : 
              item.gender === 'w' ? 'женский ' : 
              item.gender ? item.gender : 'неизвестно '}
              | {item.birthday}</p>
            </div>
            <h2 className="item-description">{item.city}</h2>
          </div>
        </>
      )}
    />
  );
};

export default EntrepreneurItem;