import React from 'react';
import { ActivityField } from './types';
import ListItem from '../List/Item';

const ActivityFieldItem: React.FC<{ activityField: ActivityField }> = ({ activityField }) => {
  return (
    <ListItem
      item={activityField}
      renderContent={(item) => (
        <>
          <div className="item-details">
            <div className="item-info">
              <h3>{item.Name}</h3>
              <p>{item.Cost}</p>
            </div>
            <p className="item-description">{item.Description}</p>
          </div>
        </>
      )}
    />
  );
};

export default ActivityFieldItem;