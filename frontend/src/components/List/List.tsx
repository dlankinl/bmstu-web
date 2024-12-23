// import React, { useState, useEffect } from 'react';
// import Pagination from '../Pagination/Pagination';
// import './List.css';

// interface ListComponentProps<T> {
//   data: T[];
//   renderItem: (item: T) => JSX.Element;
//   itemsPerPage: number;
// }

// const ListComponent = <T,>({ data, renderItem, itemsPerPage }: ListComponentProps<T>) => {
//   const [currentPage, setCurrentPage] = useState(1);
  
//   const indexOfLastItem = currentPage * itemsPerPage;
//   const indexOfFirstItem = indexOfLastItem - itemsPerPage;
//   const currentItems = data.slice(indexOfFirstItem, indexOfLastItem);

//   return (
//     <div className="list-component">
//       <div className="items">
//         {currentItems.map(renderItem)}
//       </div>
//       <Pagination
//         currentPage={currentPage}
//         totalPages={Math.ceil(data.length / itemsPerPage)}
//         onPageChange={setCurrentPage}
//       />
//     </div>
//   );
// };

// export default ListComponent;




import React, { useState } from 'react';
import Pagination from '../Pagination/Pagination';

interface ListComponentProps<T> {
  data: T[];
  renderItem: (item: T) => React.ReactNode;
  itemsPerPage: number;
  currentPage: number; // Current page prop
  totalPages: number; // Total pages prop
  onPageChange?: (pageNumber: number) => void; // Callback for when the page changes
}

const ListComponent = <T,>({ data, renderItem, itemsPerPage, currentPage, totalPages, onPageChange }: ListComponentProps<T>) => {
  
  const handlePageChange = (pageNumber: number) => {
    if (onPageChange) {
      onPageChange(pageNumber); // Notify parent about the page change
    }
  };

  return (
    <div className="list-component">
      <div className="items">
        {data.map(renderItem)}
      </div>
      <Pagination
        currentPage={currentPage}
        totalPages={totalPages} // Use total pages passed from parent
        onPageChange={handlePageChange} // Pass down the handler
      />
    </div>
  );
};

export default ListComponent;
