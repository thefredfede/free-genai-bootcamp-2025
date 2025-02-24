import React from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import StudyActivities from './pages/StudyActivities';
import StudyActivity from './pages/StudyActivity';
import Word from './pages/Word';
import Words from './pages/Words';
import Groups from './pages/Groups';
import Group from './pages/Group';
import Sessions from './pages/Sessions';
import Settings from './pages/Settings';
import Navbar from './components/Navbar';

import './App.css';

const App: React.FC = () => {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Navigate to="/dashboard" />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/study-activities" element={<StudyActivities />} />
        <Route path="/study-activities/:id" element={<StudyActivity />} />
        <Route path="/word" element={<Word />} />
        <Route path="/words/:id" element={<Words />} />
        <Route path="/groups" element={<Groups />} />
        <Route path="/groups/:id" element={<Group />} />
        <Route path="/sessions" element={<Sessions />} />
        <Route path="/settings" element={<Settings />} />
      </Routes>
    </Router>
  );
};

export default App;
