import React, { useState, useEffect } from 'react';
import { User, Edit, Save, X, Camera, Award, Target, TrendingUp } from 'lucide-react';

const Perfil: React.FC = () => {
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    name: 'Alex Johnson',
    email: 'alex.johnson@example.com',
    age: 28,
    height: 175,
    weight: 72.5,
    fitnessLevel: 'intermediate',
    goals: 'weight_loss'
  });

  useEffect(() => {
    // IMPORTANT: change id from hardcoded to fetched (auth)
    const userId = "1";

    fetch(`http://localhost:8080/users/${userId}`)
      .then(response => response.json())
      .then(data => {
        // Currently only name and email
        setFormData(prevData => ({
          ...prevData, // Keep the other fields like age, height, etc.
          name: data.name,
          email: data.email,
        }));
      })
      .catch(error => {
        console.error("Error fetching user data:", error);
      });
  }, []);

  const stats = [
    { label: 'Total Workouts', value: '156', icon: Target, color: 'emerald' },
    { label: 'Calories Burned', value: '43,250', icon: TrendingUp, color: 'orange' },
    { label: 'Achievements', value: '23', icon: Award, color: 'blue' },
    { label: 'Current Streak', value: '12 days', icon: Target, color: 'purple' }
  ];

  const achievements = [
    { title: 'First Workout', description: 'Completed your first workout', date: '2023-01-15', earned: true },
    { title: '7-Day Streak', description: 'Worked out for 7 consecutive days', date: '2023-02-20', earned: true },
    { title: '100 Workouts', description: 'Completed 100 total workouts', date: '2023-06-10', earned: true },
    { title: 'Marathon Ready', description: 'Ran 42km in a month', date: 'Not earned', earned: false },
    { title: 'Strength Master', description: 'Complete 50 strength workouts', date: '2023-08-15', earned: true },
    { title: '30-Day Challenge', description: 'Work out for 30 consecutive days', date: 'Not earned', earned: false }
  ];

  const handleSave = () => {
    setIsEditing(false);
    // Here you would typically send the data to your Go backend
    console.log('Saving profile data:', formData);
  };

  const handleCancel = () => {
    setIsEditing(false);
    // Reset form data if needed
  };

  return (
    <div className="space-y-8">
      {/* Header */}
      <div>
        <h2 className="text-2xl font-bold text-slate-900">Perfil</h2>
        <p className="text-slate-600">Manage your account and track your progress</p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Perfil Information */}
        <div className="lg:col-span-2 space-y-6">
          {/* Basic Info Card */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
            <div className="flex justify-between items-center mb-6">
              <h3 className="text-lg font-semibold text-slate-900">Personal Information</h3>
              {!isEditing ? (
                <button
                  onClick={() => setIsEditing(true)}
                  className="flex items-center space-x-2 px-4 py-2 bg-emerald-500 text-white rounded-xl hover:bg-emerald-600 transition-colors"
                >
                  <Edit className="h-4 w-4" />
                  <span>Edit</span>
                </button>
              ) : (
                <div className="flex space-x-2">
                  <button
                    onClick={handleSave}
                    className="flex items-center space-x-2 px-4 py-2 bg-emerald-500 text-white rounded-xl hover:bg-emerald-600 transition-colors"
                  >
                    <Save className="h-4 w-4" />
                    <span>Save</span>
                  </button>
                  <button
                    onClick={handleCancel}
                    className="flex items-center space-x-2 px-4 py-2 bg-slate-200 text-slate-600 rounded-xl hover:bg-slate-300 transition-colors"
                  >
                    <X className="h-4 w-4" />
                    <span>Cancel</span>
                  </button>
                </div>
              )}
            </div>

            <div className="flex items-center space-x-6 mb-6">
              <div className="relative">
                <div className="w-24 h-24 bg-gradient-to-br from-emerald-400 to-emerald-500 rounded-full flex items-center justify-center">
                  <User className="h-12 w-12 text-white" />
                </div>
                {isEditing && (
                  <button className="absolute bottom-0 right-0 bg-white p-2 rounded-full shadow-lg border border-slate-200">
                    <Camera className="h-4 w-4 text-slate-600" />
                  </button>
                )}
              </div>
              <div>
                <h2 className="text-2xl font-bold text-slate-900">{formData.name}</h2>
                <p className="text-slate-600">{formData.email}</p>
              </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Full Name</label>
                {isEditing ? (
                  <input
                    type="text"
                    value={formData.name}
                    onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  />
                ) : (
                  <p className="text-slate-900 font-medium">{formData.name}</p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Email</label>
                {isEditing ? (
                  <input
                    type="email"
                    value={formData.email}
                    onChange={(e) => setFormData({ ...formData, email: e.target.value })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  />
                ) : (
                  <p className="text-slate-900 font-medium">{formData.email}</p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Age</label>
                {isEditing ? (
                  <input
                    type="number"
                    value={formData.age}
                    onChange={(e) => setFormData({ ...formData, age: parseInt(e.target.value) })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  />
                ) : (
                  <p className="text-slate-900 font-medium">{formData.age} years</p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Height</label>
                {isEditing ? (
                  <input
                    type="number"
                    value={formData.height}
                    onChange={(e) => setFormData({ ...formData, height: parseInt(e.target.value) })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  />
                ) : (
                  <p className="text-slate-900 font-medium">{formData.height} cm</p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Weight</label>
                {isEditing ? (
                  <input
                    type="number"
                    step="0.1"
                    value={formData.weight}
                    onChange={(e) => setFormData({ ...formData, weight: parseFloat(e.target.value) })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  />
                ) : (
                  <p className="text-slate-900 font-medium">{formData.weight} kg</p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">Fitness Level</label>
                {isEditing ? (
                  <select
                    value={formData.fitnessLevel}
                    onChange={(e) => setFormData({ ...formData, fitnessLevel: e.target.value })}
                    className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500"
                  >
                    <option value="beginner">Beginner</option>
                    <option value="intermediate">Intermediate</option>
                    <option value="advanced">Advanced</option>
                  </select>
                ) : (
                  <p className="text-slate-900 font-medium capitalize">{formData.fitnessLevel}</p>
                )}
              </div>
            </div>
          </div>

          {/* Achievements */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
            <h3 className="text-lg font-semibold text-slate-900 mb-6">Achievements</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {achievements.map((achievement, index) => (
                <div
                  key={index}
                  className={`p-4 rounded-xl border-2 transition-all ${
                    achievement.earned
                      ? 'border-emerald-200 bg-emerald-50'
                      : 'border-slate-200 bg-slate-50 opacity-60'
                  }`}
                >
                  <div className="flex items-center space-x-3">
                    <div className={`p-2 rounded-lg ${achievement.earned ? 'bg-emerald-500' : 'bg-slate-400'}`}>
                      <Award className="h-5 w-5 text-white" />
                    </div>
                    <div className="flex-1">
                      <h4 className="font-medium text-slate-900">{achievement.title}</h4>
                      <p className="text-sm text-slate-600">{achievement.description}</p>
                      <p className="text-xs text-slate-400">{achievement.date}</p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* Stats Sidebar */}
        <div className="space-y-6">
          {/* Quick Stats */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
            <h3 className="text-lg font-semibold text-slate-900 mb-6">Your Stats</h3>
            <div className="space-y-4">
              {stats.map((stat, index) => {
                const Icon = stat.icon;
                return (
                  <div key={index} className="flex items-center space-x-4">
                    <div className={`p-3 rounded-xl ${
                      stat.color === 'emerald' ? 'bg-emerald-100' :
                      stat.color === 'orange' ? 'bg-orange-100' :
                      stat.color === 'blue' ? 'bg-blue-100' : 'bg-purple-100'
                    }`}>
                      <Icon className={`h-5 w-5 ${
                        stat.color === 'emerald' ? 'text-emerald-600' :
                        stat.color === 'orange' ? 'text-orange-600' :
                        stat.color === 'blue' ? 'text-blue-600' : 'text-purple-600'
                      }`} />
                    </div>
                    <div className="flex-1">
                      <p className="text-sm text-slate-600">{stat.label}</p>
                      <p className="text-lg font-semibold text-slate-900">{stat.value}</p>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          {/* BMI Calculator */}
          <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
            <h3 className="text-lg font-semibold text-slate-900 mb-4">BMI Calculator</h3>
            <div className="text-center">
              <div className="text-3xl font-bold text-emerald-600 mb-2">
                {(formData.weight / Math.pow(formData.height / 100, 2)).toFixed(1)}
              </div>
              <div className="text-sm text-slate-600 mb-4">Normal Weight</div>
              <div className="w-full bg-slate-200 rounded-full h-2">
                <div className="bg-emerald-500 h-2 rounded-full" style={{ width: '60%' }} />
              </div>
              <div className="flex justify-between text-xs text-slate-500 mt-2">
                <span>Underweight</span>
                <span>Normal</span>
                <span>Overweight</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Perfil;