import React, { useState, useEffect } from 'react';
import Pagination from '../Pagination/Pagination';
import './List.css';

interface ListComponentProps<T> {
  data: T[];
  renderItem: (item: T) => JSX.Element;
  itemsPerPage: number;
}

const ListComponent = <T,>({ data, renderItem, itemsPerPage }: ListComponentProps<T>) => {
  const [currentPage, setCurrentPage] = useState(1);
  
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = data.slice(indexOfFirstItem, indexOfLastItem);

  return (
    <div className="list-component">
      <div className="items">
        {currentItems.map(renderItem)}
      </div>
      <Pagination
        currentPage={currentPage}
        totalPages={Math.ceil(data.length / itemsPerPage)}
        onPageChange={setCurrentPage}
      />
    </div>
  );
};

export default ListComponent;
