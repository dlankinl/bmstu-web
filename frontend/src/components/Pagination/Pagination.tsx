import React from 'react';
import './Pagination.css';

interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

const Pagination: React.FC<PaginationProps> = ({ currentPage, totalPages, onPageChange }) => {
  const handlePrevious = () => {
    if (currentPage > 1) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNext = () => {
    if (currentPage < totalPages) {
      onPageChange(currentPage + 1);
    }
  };

  return (
    <div className="pagination">
      <button
        onClick={handlePrevious}
        className={`arrow-button ${currentPage === 1 ? 'disabled' : ''}`}
        disabled={currentPage === 1}
      >
        &lt; {/* Left Arrow */}
      </button>
      <div className="page-numbers">
        {Array.from({ length: totalPages }, (_, index) => (
          <button
            key={index + 1}
            onClick={() => onPageChange(index + 1)}
            className={`page-item ${currentPage === index + 1 ? 'active' : ''}`}
          >
            {index + 1}
          </button>
        ))}
      </div>
      <button
        onClick={handleNext}
        className={`arrow-button ${currentPage === totalPages ? 'disabled' : ''}`}
        disabled={currentPage === totalPages}
      >
        &gt; {/* Right Arrow */}
      </button>
    </div>
  );
};

export default Pagination; // Make sure to export the component
