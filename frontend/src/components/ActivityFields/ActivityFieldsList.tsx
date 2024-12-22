import React, { useEffect, useState } from 'react';
import ListComponent from '../List/List';
import ActivityFieldItem from './ActivityFieldItem';
import { ActivityField } from './types';
import ListContainer from '../List/ListContainer';
import GreenButton from '../Buttons/GreenButton';
import Modal from '../Modal/Modal';
import "./AcitivityFieldsList.css";
import CreateActivityFieldComponent from '../CreateActivityField/CreateActivityField';

interface ActivityFieldsListProps {
  isAdmin: boolean;
}

const ActivityFieldsList: React.FC<ActivityFieldsListProps> = ({ isAdmin }) => {
  const [activityFields, setActivityFields] = useState<ActivityField[]>([]);
  const [isModalActive, setModalActive] = useState(false);
  
  const handleModalOpen = () => {
    setModalActive(true);
  };
  const handleModalClose = () => {
    setModalActive(false);
  };

  useEffect(() => {
    const fetchActivityFields = () => {
      const staticActivityFields: ActivityField[] = [
        { ID: "1", Name: "Test1", Cost: 1.0, Description: "Small Test1 AF" },
        { ID: "2", Name: "Test2", Cost: 2.0, Description: "Small Test2 AF" },
        { ID: "3", Name: "Test3", Cost: 3.0, Description: "Small Test3 AF" },
        { ID: "4", Name: "Test4", Cost: 4.0, Description: "Small Test4 AF" },
        { ID: "5", Name: "Test5", Cost: 5.0, Description: "Small Test5 AF" },
        { ID: "6", Name: "Test6", Cost: 6.0, Description: "Small Test6 AF" },
      ];
      setActivityFields(staticActivityFields);
    };

    fetchActivityFields();
  }, []);

  return (
    <ListContainer>
      <div style={{ position: 'relative' }}>
      {isAdmin && (
        <div className="button-container">
          <GreenButton text={"Добавить"} icon={""} onClick={handleModalOpen}/>
        </div>
      )}
        <ListComponent
          data={activityFields}
          renderItem={(activityField) => (
            <ActivityFieldItem key={activityField.ID} activityField={activityField} />
          )}
          itemsPerPage={3}
        />
        <div>
          {isModalActive && (
            <Modal onClose={handleModalClose}>
              <CreateActivityFieldComponent/>
            </Modal>
          )}
        </div>
      </div>
    </ListContainer>
  );
};

export default ActivityFieldsList;