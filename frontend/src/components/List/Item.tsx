import React from 'react';
import './Item.css'

interface ListItemProps<T> {
  item: T; 
  renderContent: (item: T) => JSX.Element;
}

const ListItem = <T,>({ item, renderContent }: ListItemProps<T>) => {
  return (
    <div className="list-item">
      {renderContent(item)} {/* Render the content using the provided function */}
    </div>
  );
};

export default ListItem;