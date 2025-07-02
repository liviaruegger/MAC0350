import React, { useState } from 'react';
import { Save, Clock, Waves, MapPin, Thermometer, Eye, Heart, Calendar } from 'lucide-react';

interface Interval {
    distance: number;
    type: string;
    stroke: string;
    time: string;
    rest: number;
    notes?: string;
}

interface Activity {
    date: string;
    pool: string;
    poolLength: number;
    duration: number; // in minutes
    distance: number; // in meters
    strokes: {
        freestyle?: number;
        backstroke?: number;
        breaststroke?: number;
        butterfly?: number;
    };
    intervals: Interval[];
    waterTemp?: number;
    notes: string;
    feeling: 'excellent' | 'good' | 'average' | 'tired' | 'poor';
    heartRate?: {
        avg: number;
        max: number;
    };
}

const AtividadeLog: React.FC = () => {
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [submitMessage, setSubmitMessage] = useState<string | null>(null);
    
    const [activity, setActivity] = useState<Activity>({
        date: new Date().toISOString().split('T')[0],
        pool: '',
        poolLength: 25,
        duration: 0,
        distance: 0,
        strokes: {},
        intervals: [],
        notes: '',
        feeling: 'good',
        heartRate: { avg: 0, max: 0 }
    });

    const [currentInterval, setCurrentInterval] = useState({
        distance: 0,
        type: 'swim',
        stroke: 'freestyle',
        time: '',
        rest: 0,
        notes: ''
    });

    const typeOptions = [
        { value: 'swim', label: 'Swim' },
        { value: 'rest', label: 'Rest' },
        { value: 'drill', label: 'Drill' },
        { value: 'kick', label: 'Kick' },
        { value: 'pull', label: 'Pull' },
        { value: 'warmup', label: 'Warmup' },
        { value: 'main_set', label: 'Main Set' },
        { value: 'cooldown', label: 'Cooldown' }
    ];

    const strokeOptions = [
        { value: 'freestyle', label: 'Freestyle' },
        { value: 'backstroke', label: 'Backstroke' },
        { value: 'breaststroke', label: 'Breaststroke' },
        { value: 'butterfly', label: 'Butterfly' },
        { value: 'medley', label: 'Individual Medley' },
        { value: 'kick', label: 'Kick Only' },
        { value: 'drill', label: 'Drill' }
    ];

    const feelingOptions = [
        { value: 'excellent', label: 'Excellent', color: 'text-green-600', bg: 'bg-green-100' },
        { value: 'good', label: 'Good', color: 'text-blue-600', bg: 'bg-blue-100' },
        { value: 'average', label: 'Average', color: 'text-yellow-600', bg: 'bg-yellow-100' },
        { value: 'tired', label: 'Tired', color: 'text-orange-600', bg: 'bg-orange-100' },
        { value: 'poor', label: 'Poor', color: 'text-red-600', bg: 'bg-red-100' }
    ];

    const addInterval = () => {
        if (currentInterval.distance > 0 && currentInterval.time) {
            const newInterval: Interval = {
                ...currentInterval
            }

            setActivity(prev => ({
                ...prev,
                intervals: [...prev.intervals, newInterval]
            }));
            setCurrentInterval({
                distance: 0,
                type: 'swim',
                stroke: 'freestyle',
                time: '',
                rest: 0,
                notes: ''
            });
        }
    };

    const removeInterval = (index: number) => {
        setActivity(prev => ({
            ...prev,
            intervals: prev.intervals.filter((_, i) => i !== index)
        }));
    };

    const calculateTotalDistance = () => {
        return activity.intervals.reduce((total, interval) => total + interval.distance, 0);
    };

    const calculatePace = () => {
        const totalDistance = calculateTotalDistance();
        if (totalDistance === 0 || activity.duration === 0) return '0:00';
        
        const pacePerHundred = (activity.duration * 60) / (totalDistance / 100);
        const minutes = Math.floor(pacePerHundred / 60);
        const seconds = Math.floor(pacePerHundred % 60);
        return `${minutes}:${seconds.toString().padStart(2, '0')}`;
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setIsSubmitting(true);
        setSubmitMessage(null);

        try {
            // Simulate API call delay
            await new Promise(resolve => setTimeout(resolve, 1000));

            // Calculate final distance if not manually entered
            const finalActivity = {
                ...activity,
                distance: activity.distance || calculateTotalDistance()
            };

            // Simulate successful save (in real app, this would be an API call)
            console.log('Saving activity:', finalActivity);
            
            setSubmitMessage('Atividade salva com sucesso!');
            
            // Reset form after successful submission
            setTimeout(() => {
                setActivity({
                    date: new Date().toISOString().split('T')[0],
                    pool: '',
                    poolLength: 25,
                    duration: 0,
                    distance: 0,
                    strokes: {},
                    intervals: [],
                    notes: '',
                    feeling: 'good',
                    heartRate: { avg: 0, max: 0 }
                });
                setSubmitMessage(null);
            }, 2000);

        } catch (error) {
            console.error('Error saving activity:', error);
            setSubmitMessage('Failed to save activity');
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="space-y-6">
            {/* Header */}
            <div>
                <h2 className="text-2xl font-bold text-slate-900">Adicionar Atividade</h2>
                <p className="text-slate-600">Registre os detalhes da sua sessão de natação</p>
            </div>

            {/* Success/Error Message */}
            {submitMessage && (
                <div className={`p-4 rounded-xl ${
                    submitMessage.includes('successfully') 
                        ? 'bg-green-50 border border-green-200 text-green-800' 
                        : 'bg-red-50 border border-red-200 text-red-800'
                }`}>
                    {submitMessage}
                </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-6">
                {/* Basic Information */}
                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <h3 className="text-lg font-semibold text-slate-900 mb-4 flex items-center space-x-2">
                        <Calendar className="h-5 w-5 text-blue-500" />
                        <span>Basic Information</span>
                    </h3>
                    
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Data</label>
                            <input
                                type="date"
                                value={activity.date}
                                onChange={(e) => setActivity(prev => ({ ...prev, date: e.target.value }))}
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                required
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Piscina/Local</label>
                            <input
                                type="text"
                                value={activity.pool}
                                onChange={(e) => setActivity(prev => ({ ...prev, pool: e.target.value }))}
                                placeholder="ex.: CEPE-USP"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                required
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Comprimento (m)</label>
                            <select
                                value={activity.poolLength}
                                onChange={(e) => setActivity(prev => ({ ...prev, poolLength: parseInt(e.target.value) }))}
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            >
                                <option value={25}>25m</option>
                                <option value={50}>50m</option>
                                <option value={33}>33.3m (jardas)</option>
                            </select>
                        </div>
                    </div>
                </div>

                {/* Workout Details */}
                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <h3 className="text-lg font-semibold text-slate-900 mb-4 flex items-center space-x-2">
                        <Waves className="h-5 w-5 text-blue-500" />
                        <span>Detalhes da Atividade</span>
                    </h3>
                    
                    <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6">
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Duração (minutos)</label>
                            <input
                                type="number"
                                value={activity.duration || ''}
                                onChange={(e) => setActivity(prev => ({ ...prev, duration: parseInt(e.target.value) || 0 }))}
                                placeholder="60"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                required
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Distância Total (m)</label>
                            <input
                                type="number"
                                value={activity.distance || calculateTotalDistance() || ''}
                                onChange={(e) => setActivity(prev => ({ ...prev, distance: parseInt(e.target.value) || 0 }))}
                                placeholder="Auto-calculated from sets"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Temperatura da Água (°C)</label>
                            <input
                                type="number"
                                value={activity.waterTemp || ''}
                                onChange={(e) => setActivity(prev => ({ ...prev, waterTemp: parseInt(e.target.value) || undefined }))}
                                placeholder="26"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Como você se sentiu?</label>
                            <select
                                value={activity.feeling}
                                onChange={(e) => setActivity(prev => ({ ...prev, feeling: e.target.value as any }))}
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            >
                                {feelingOptions.map(option => (
                                    <option key={option.value} value={option.value}>{option.label}</option>
                                ))}
                            </select>
                        </div>
                    </div>

                    {/* Heart Rate */}
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Frequência Cardíaca Média</label>
                            <input
                                type="number"
                                value={activity.heartRate?.avg || ''}
                                onChange={(e) => setActivity(prev => ({ 
                                    ...prev, 
                                    heartRate: { ...prev.heartRate!, avg: parseInt(e.target.value) || 0 }
                                }))}
                                placeholder="140"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            />
                        </div>
                        
                        <div>
                            <label className="block text-sm font-medium text-slate-700 mb-2">Frequência Cardíaca Máxima</label>
                            <input
                                type="number"
                                value={activity.heartRate?.max || ''}
                                onChange={(e) => setActivity(prev => ({ 
                                    ...prev, 
                                    heartRate: { ...prev.heartRate!, max: parseInt(e.target.value) || 0 }
                                }))}
                                placeholder="165"
                                className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                            />
                        </div>
                    </div>
                </div>

                {/* Sets */}
                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <h3 className="text-lg font-semibold text-slate-900 mb-4">Intervalos</h3>
                    
                    {/* Add Set Form */}
                    <div className="bg-slate-50 rounded-xl p-4 mb-4">
                        <h4 className="font-medium text-slate-900 mb-3">Adicionar Intervalo</h4>
                        <div className="grid grid-cols-1 md:grid-cols-3 gap-3 mb-3">

                            <div>
                                <label className="block text-sm font-medium text-slate-700 mb-1">Distância (m)</label>
                                <input
                                    type="number"
                                    value={currentInterval.distance || ''}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, distance: parseInt(e.target.value) || 0 }))}
                                    placeholder="100"
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                />
                            </div>

                            <div>
                                <label className="block text-sm font-medium text-slate-700 mb-1">Tipo</label>
                                <select
                                    value={currentInterval.type}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, stroke: e.target.value }))}
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                >
                                    {typeOptions.map(option => (
                                        <option key={option.value} value={option.value}>{option.label}</option>
                                    ))}
                                </select>
                            </div>
                            
                            <div>
                                <label className="block text-sm font-medium text-slate-700 mb-1">Estilo</label>
                                <select
                                    value={currentInterval.stroke}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, stroke: e.target.value }))}
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                >
                                    {strokeOptions.map(option => (
                                        <option key={option.value} value={option.value}>{option.label}</option>
                                    ))}
                                </select>
                            </div>
                            
                            <div>
                                <label className="block text-sm font-medium text-slate-700 mb-1">Tempo</label>
                                <input
                                    type="text"
                                    value={currentInterval.time}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, time: e.target.value }))}
                                    placeholder="1:30.00"
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                />
                            </div>

                            <div>
                                <label className="block text-sm font-medium text-slate-700 mb-1">Descanso (segundos)</label>
                                <input
                                    type="number"
                                    value={currentInterval.rest || ''}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, rest: parseInt(e.target.value) || 0 }))}
                                    placeholder="30"
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                />
                            </div>
                        </div>
                        
                        <div className="grid grid-cols-1 md:grid-cols-4 gap-3">
                            <div className="md:col-span-4">
                                <label className="block text-sm font-medium text-slate-700 mb-1">Notas (opcional)</label>
                                <input
                                    type="text"
                                    value={currentInterval.notes}
                                    onChange={(e) => setCurrentInterval(prev => ({ ...prev, notes: e.target.value }))}
                                    placeholder="ex.: foquei na pegada, senti força, respiração bilateral"
                                    className="w-full px-3 py-2 rounded-lg border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                                />
                            </div>
                            
                            <div className="flex items-end">
                                <button
                                    type="button"
                                    onClick={addInterval}
                                    className="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 transition-colors"
                                >
                                    Adicionar Intervalo
                                </button>
                            </div>
                        </div>
                    </div>

                    {/* Sets List */}
                    {activity.intervals.length > 0 && (
                        <div className="space-y-3">
                            <h4 className="font-medium text-slate-900">Intervalos ({activity.intervals.length})</h4>
                            {activity.intervals.map((interval, index) => (
                                <div key={index} className="p-4 bg-blue-50 rounded-lg border border-blue-100">
                                    <div className="flex items-center justify-between mb-2">
                                        <div className="flex items-center space-x-4">
                                            <span className="font-medium text-blue-900">{interval.distance}m</span>
                                            <span className="text-slate-600 capitalize">{interval.stroke}</span>
                                            <span className="text-slate-600">{interval.time}</span>
                                            {interval.rest > 0 && <span className="text-slate-500">Descanso: {interval.rest}s</span>}
                                        </div>
                                        <button
                                            type="button"
                                            onClick={() => removeInterval(index)}
                                            className="text-red-500 hover:text-red-700 text-sm font-medium"
                                        >
                                            Remover
                                        </button>
                                    </div>
                                    {interval.notes && (
                                        <div className="text-sm text-slate-600 bg-white p-2 rounded border border-slate-200">
                                            <span className="font-medium">Notas:</span> {interval.notes}
                                        </div>
                                    )}
                                </div>
                            ))}
                            
                            {/* Summary */}
                            <div className="mt-4 p-3 bg-blue-100 rounded-lg">
                                <div className="flex justify-between text-sm">
                                    <span>Distância Total: <strong>{calculateTotalDistance()}m</strong></span>
                                    <span>Ritmo Médio: <strong>{calculatePace()}/100m</strong></span>
                                </div>
                            </div>
                        </div>
                    )}
                </div>

                {/* Notes */}
                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <h3 className="text-lg font-semibold text-slate-900 mb-4">Notas Gerais</h3>
                    <textarea
                        value={activity.notes}
                        onChange={(e) => setActivity(prev => ({ ...prev, notes: e.target.value }))}
                        placeholder="Como foi sua natação? Alguma observação, foco na técnica ou metas para a próxima vez?"
                        rows={4}
                        className="w-full px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    />
                </div>

                {/* Submit Button */}
                <div className="flex space-x-4">
                    <button
                        type="submit"
                        disabled={isSubmitting}
                        className="flex-1 bg-gradient-to-r from-blue-500 to-cyan-500 text-white py-4 rounded-xl font-medium hover:shadow-lg hover:shadow-blue-500/25 transition-all flex items-center justify-center space-x-2 disabled:opacity-50"
                    >
                        <Save className="h-5 w-5" />
                        <span>{isSubmitting ? 'Saving...' : 'Save Swimming Workout'}</span>
                    </button>
                </div>
            </form>
        </div>
    );
};

export default AtividadeLog;