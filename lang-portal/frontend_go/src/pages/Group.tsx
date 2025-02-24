import React from 'react';
import { useParams } from 'react-router-dom';

const Group: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  return (
    <div>
      <h1>Group {id}</h1>
    </div>
  );
};

export default Group;