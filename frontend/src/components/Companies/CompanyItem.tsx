import React from 'react';
import { Company } from './types';
import ListItem from '../List/Item';
import { Link } from 'react-router-dom';

const CompanyItem: React.FC<{ company: Company }> = ({ company }) => {
  return (
    <ListItem
      item={company}
      renderContent={(item) => (
        <>
          <div className="item-details">
            <div className="item-info">
              <Link to={`/companies/${item.id}`}>
                <h3>{item.name}</h3>
              </Link>
              <p>{item.city}</p>
            </div>
            <p className="item-description">{item.description}</p>
          </div>
        </>
      )}
    />
  );
};

export default CompanyItem;