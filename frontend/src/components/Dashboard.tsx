import React from 'react';
import { Target } from 'lucide-react';

const Dashboard: React.FC = () => {
  return (
    <div className="space-y-8">

      {/* Welcome Section */}
      <div className="bg-gradient-to-r from-cyan-600 to-cyan-500 rounded-2xl p-6 text-white">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-2xl font-bold mb-2">Bom dia, Antônio!</h2>
            <p className="text-cyan-100">Você já bateu sua meta hoje</p>
          </div>
          <div className="bg-white/20 p-3 rounded-xl backdrop-blur-sm">
            <Target className="h-8 w-8" />
          </div>
        </div>
      </div>

    </div>
  );
};

export default Dashboard;