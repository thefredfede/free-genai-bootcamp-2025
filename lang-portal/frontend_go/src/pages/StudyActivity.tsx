import React from 'react';
import { useParams } from 'react-router-dom';

const StudyActivity: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  return (
    <div>
      <h1>Study Activity {id}</h1>
    </div>
  );
};

export default StudyActivity;