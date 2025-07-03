import { useState } from 'react';
import { BarChart3, Settings, User, Waves } from 'lucide-react';
import Dashboard from './components/Dashboard';
import Perfil from './components/Perfil';


function App() {
const [activeTab, setActiveTab] = useState('dashboard');

const navigation = [
	{ id: 'dashboard', label: 'Dashboard', icon: BarChart3 },
	{ id: 'perfil', label: 'Perfil', icon: User },
];

const renderContent = () => {
	switch (activeTab) {
	case 'dashboard':
		return <Dashboard />;
	case 'perfil':
		return <Perfil />;
	default:
		return <Dashboard />;
	}
};

return (
	<div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
	{/* Header */}
	<header className="bg-white shadow-sm border-b border-slate-200">
		<div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
		<div className="flex justify-between items-center h-16">
			<div className="flex items-center space-x-3">
			<div className="bg-gradient-to-r from-blue-500 to-blue-400 p-2 rounded-xl">
				<Waves className="h-6 w-6 text-white" />
			</div>
			<h1 className="text-2xl font-bold bg-gradient-to-r from-blue-700 to-cyan-600 bg-clip-text text-transparent">
				Swim Tracker
			</h1>
			</div>
			<button className="p-2 rounded-lg hover:bg-slate-100 transition-colors">
			<Settings className="h-5 w-5 text-slate-600" />
			</button>
		</div>
		</div>
	</header>

	<div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
		<div className="flex flex-col lg:flex-row gap-8">
		{/* Sidebar Navigation */}
		<nav className="lg:w-64 space-y-2">
			{navigation.map((item) => {
			const Icon = item.icon;
			return (
				<button
				key={item.id}
				onClick={() => setActiveTab(item.id)}
				className={`w-full flex items-center space-x-3 px-4 py-3 rounded-xl font-medium transition-all duration-200 ${
					activeTab === item.id
					? 'bg-gradient-to-r from-blue-500 to-cyan-400 text-white shadow-lg shadow-blue-500/25'
					: 'text-slate-600 hover:bg-white hover:shadow-sm'
				}`}
				>
				<Icon className="h-5 w-5" />
				<span>{item.label}</span>
				</button>
			);
			})}
		</nav>

		{/* Main Content */}
		<main className="flex-1 min-h-screen">
			{renderContent()}
		</main>
		</div>
	</div>
	</div>
);
}

export default App;