import React, { useState, useEffect } from 'react';
import { User, Edit, Save, X, Camera, Award, Target, TrendingUp } from 'lucide-react';

const Perfil: React.FC = () => {
const [isEditing, setIsEditing] = useState(false);
const [formData, setFormData] = useState({
	name: 'José Exemplar',
	email: 'jose@examplo.com',
	city: 'Exemplópolis',
	phone: '(00) 0 0000-0000',
	age: 27,
	height: 180,
	weight: 75.5,
	goals: 'weight_loss'
});

useEffect(() => {
	// IMPORTANT: change id from hardcoded to fetched (auth)
	const userId = "9cdba1c6-9a50-464f-a892-3efd75090243";

	fetch(`http://localhost:8080/users/${userId}`)
	.then(response => response.json())
	.then(data => {
		// Currently only name and email
		setFormData(prevData => ({
		...prevData, // Keep the other fields like age, height, etc.
		name: data.name,
		email: data.email,
		city: data.city,
		phone: data.phone,
		}));
	})
	.catch(error => {
		console.error("Error fetching user data:", error);
	});
}, []);

const stats = [
	{ label: 'Atividades totais', value: '56', icon: Target, color: 'cyan' },
	{ label: 'Distância total percorrida', value: '213,7 km', icon: TrendingUp, color: 'orange' },
	{ label: 'Conquistas', value: '23', icon: Award, color: 'blue' },
	{ label: 'Sequência atual', value: '15 dias', icon: Target, color: 'purple' }
];

const achievements = [
	{ title: 'Primeira atividade', description: 'Você completou sua primeira atividade', date: '2023-01-15', earned: true },
	{ title: 'Sequência de 7 Dias', description: 'Cumpriu 7 dias consecutivos', date: '2023-02-20', earned: true },
	{ title: '100 Atividades', description: 'Completou um total de 100 atividades', date: '2023-06-10', earned: true },
	{ title: 'Pronto para a Maratona!', description: 'Nadou 10km em uma sessão', date: 'Não conquistado', earned: false },
	{ title: 'Desafio de 30 Dias', description: 'Cumpriu 30 dias consecutivos', date: 'Não conquistado', earned: false }
];

const handleSave = () => {
	const userId = "9cdba1c6-9a50-464f-a892-3efd75090243";

	const userToUpdate = {
		name: formData.name,
		email: formData.email,
		// Your form doesn't have city and phone yet, so we'll send empty strings for now.
		// We can add them to the form later.
		city: formData.city,
		phone: formData.phone
		};

		fetch(`http://localhost:8080/users/${userId}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(userToUpdate),
		})
		.then(response => {
			if (!response.ok) {
				throw new Error('Failed to update user');
			}
			return response.json();
		})
		.then(updatedUser => {
			console.log('Profile updated successfully:', updatedUser);
			setIsEditing(false); // Exit editing mode only on success
		})
		.catch(error => {
			console.error('Error updating user data:', error);
		});
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
		<p className="text-slate-600">Gerencie sua conta e acompanhe seu progresso.</p>
	</div>

	<div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
		{/* Perfil Information */}
		<div className="lg:col-span-2 space-y-6">
		{/* Basic Info Card */}
		<div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
			<div className="flex justify-between items-center mb-6">
			<h3 className="text-lg font-semibold text-slate-900">Dados Pessoais</h3>
			{!isEditing ? (
				<button
				onClick={() => setIsEditing(true)}
				className="flex items-center space-x-2 px-4 py-2 bg-blue-400 text-white rounded-xl hover:bg-cyan-500 transition-colors"
				>
				<Edit className="h-4 w-4" />
				<span>Editar</span>
				</button>
			) : (
				<div className="flex space-x-2">
				<button
					onClick={handleSave}
					className="flex items-center space-x-2 px-4 py-2 bg-cyan-500 text-white rounded-xl hover:bg-cyan-600 transition-colors"
				>
					<Save className="h-4 w-4" />
					<span>Salvar</span>
				</button>
				<button
					onClick={handleCancel}
					className="flex items-center space-x-2 px-4 py-2 bg-slate-200 text-slate-600 rounded-xl hover:bg-slate-300 transition-colors"
				>
					<X className="h-4 w-4" />
					<span>Cancelar</span>
				</button>
				</div>
			)}
			</div>

			<div className="flex items-center space-x-6 mb-6">
			<div className="relative">
				<div className="w-24 h-24 bg-gradient-to-br from-blue-400 to-cyan-400 rounded-full flex items-center justify-center">
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
				<label className="block text-sm font-medium text-slate-700 mb-2">Nome Completo</label>
				{isEditing ? (
				<input
					type="text"
					value={formData.name}
					onChange={(e) => setFormData({ ...formData, name: e.target.value })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
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
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.email}</p>
				)}
			</div>

			<div>
				<label className="block text-sm font-medium text-slate-700 mb-2">Idade</label>
				{isEditing ? (
				<input
					type="number"
					value={formData.age}
					onChange={(e) => setFormData({ ...formData, age: parseInt(e.target.value) })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.age} anos</p>
				)}
			</div>

			<div>
				<label className="block text-sm font-medium text-slate-700 mb-2">Altura</label>
				{isEditing ? (
				<input
					type="number"
					value={formData.height}
					onChange={(e) => setFormData({ ...formData, height: parseInt(e.target.value) })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.height} cm</p>
				)}
			</div>

			<div>
				<label className="block text-sm font-medium text-slate-700 mb-2">Peso</label>
				{isEditing ? (
				<input
					type="number"
					value={formData.weight}
					onChange={(e) => setFormData({ ...formData, weight: parseFloat(e.target.value) })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.weight} kg</p>
				)}
			</div>

			<div>
				<label className="block text-sm font-medium text-slate-700 mb-2">Cidade</label>
				{isEditing ? (
				<input
					type="text"
					value={formData.weight}
					onChange={(e) => setFormData({ ...formData, city: e.target.value })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.city}</p>
				)}
			</div>

			
			<div>
				<label className="block text-sm font-medium text-slate-700 mb-2">Telefone</label>
				{isEditing ? (
				<input
					type="text"
					value={formData.phone}
					onChange={(e) => setFormData({ ...formData, city: e.target.value })}
					className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-cyan-500 focus:border-cyan-500"
				/>
				) : (
				<p className="text-slate-900 font-medium">{formData.phone}</p>
				)}
			</div>

			</div>
		</div>

		<div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
			<h3 className="text-lg font-semibold text-slate-900 mb-6">Conquistas</h3>
			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
			{achievements.map((achievement, index) => (
				<div
				key={index}
				className={`p-4 rounded-xl border-2 transition-all ${
					achievement.earned
					? 'border-cyan-200 bg-cyan-50'
					: 'border-slate-200 bg-slate-50 opacity-60'
				}`}
				>
				<div className="flex items-center space-x-3">
					<div className={`p-2 rounded-lg ${achievement.earned ? 'bg-blue-400' : 'bg-slate-400'}`}>
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
				<h3 className="text-lg font-semibold text-slate-900 mb-6">Seus números</h3>
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

		</div>
	</div>
	</div>
);
};

export default Perfil;