import React, { useState, useEffect } from 'react';
import { Calendar, Clock, Waves, MapPin, Thermometer, Heart, Filter, Search, TrendingUp } from 'lucide-react';

interface Interval {
    id: string,
    distance: number;
    type: string;
    stroke: string;
    time: string;
    notes?: string;
}

interface Activity {
    id: string,
    date: string;
    locationName: string;
    locationType: string;
    poolSize: number;
    duration: number; // in minutes
    distance: number; // in meters
    intervals: Interval[];
    waterTemp?: number;
    notes: string;
    feeling: 'excellent' | 'good' | 'regular' | 'tired' | 'bad';
    heartRateAvg?: number;
    heartRateMax?: number;
}

const Historico: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [selectedFeeling, setSelectedFeeling] = useState('all');
    const [sortBy, setSortBy] = useState('date');
    const [expandedActivities, setExpandedActivities] = useState<string[]>([]);
    const [atividades, setAtividades] = useState<Activity[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    // Replace with actual user_id logic
    const userId = '59f4e428-9d42-4af8-a18d-1e1dabef47e0';

    useEffect(() => {
        const fetchActivities = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await fetch(`http://localhost:8080/users/${userId}/activities`);
                if (!response.ok) throw new Error('Erro ao buscar atividades');
                const data = await response.json();
                // Map API response to Activity[]
                const mapped: Activity[] = (data.activities || []).map((a: any) => ({
                    id: a.id,
                    date: a.date,
                    locationName: a.location_name,
                    locationType: a.location_type,
                    poolSize: a.pool_size || 0,
                    duration: parseDurationToMinutes(a.duration),
                    distance: a.distance,
                    intervals: (a.intervals || []).map((i: any) => ({
                        id: i.id,
                        distance: i.distance,
                        type: i.type,
                        stroke: i.stroke,
                        time: i.duration || '',
                        notes: i.notes
                    })),
                    waterTemp: undefined, // Not present in API
                    notes: a.notes || '',
                    feeling: a.feeling || 'regular',
                    heartRateAvg: a.heart_rate_avg,
                    heartRateMax: a.heart_rate_max
                }));
                setAtividades(mapped);
            } catch (err: any) {
                setError(err.message || 'Erro desconhecido');
            } finally {
                setLoading(false);
            }
        };
        fetchActivities();
    }, [userId]);

    const feelingOptions = [
        { value: 'all', label: 'Todas as Sensações' },
        { value: 'excellent', label: 'Excelente', color: 'text-green-600', bg: 'bg-green-100' },
        { value: 'good', label: 'Bem', color: 'text-blue-600', bg: 'bg-blue-100' },
        { value: 'regular', label: 'Regular', color: 'text-yellow-600', bg: 'bg-yellow-100' },
        { value: 'tired', label: 'Cansado', color: 'text-orange-600', bg: 'bg-orange-100' },
        { value: 'bad', label: 'Mal', color: 'text-red-600', bg: 'bg-red-100' }
    ];

    const sortOptions = [
        { value: 'date', label: 'Mais Recente' },
        { value: 'distance', label: 'Distância' },
        { value: 'duration', label: 'Duração' },
        { value: 'feeling', label: 'Sensação' }
    ];

    const getFeelingStyle = (feeling: string) => {
        const option = feelingOptions.find(opt => opt.value === feeling);
        return option ? { color: option.color, bg: option.bg } : { color: 'text-slate-600', bg: 'bg-slate-100' };
    };

    const calculatePace = (distance: number, duration: number) => {
        if (distance === 0 || duration === 0) return '0:00';
        const pacePerHundred = (duration * 60) / (distance / 100);
        const minutes = Math.floor(pacePerHundred / 60);
        const seconds = Math.floor(pacePerHundred % 60);
        return `${minutes}:${seconds.toString().padStart(2, '0')}`;
    };

    // Helper to format duration as 'Xh Ym' or 'Xm'
    const formatDuration = (minutes: number) => {
        if (minutes >= 60) {
            const h = Math.floor(minutes / 60);
            const m = minutes % 60;
            return m > 0 ? `${h}h ${m}m` : `${h}h`;
        }
        return `${minutes} min`;
    };

    // Helper to parse duration strings like '1h 15m', '75m', '2h', '60m', etc.
    function parseDurationToMinutes(duration: string | number): number {
        if (typeof duration === 'number') return duration;
        if (!duration) return 0;
        let min = 0;
        const hMatch = duration.match(/(\d+)h/);
        const mMatch = duration.match(/(\d+)m/);
        if (hMatch) min += parseInt(hMatch[1], 10) * 60;
        if (mMatch) min += parseInt(mMatch[1], 10);
        // If only a number (e.g. '75'), treat as minutes
        if (!hMatch && !mMatch && /^\d+$/.test(duration)) min = parseInt(duration, 10);
        return min;
    }

    const filteredAndSortedWorkouts = atividades
        .filter(atividades => {
            const matchesSearch = atividades.locationName.toLowerCase().includes(searchTerm.toLowerCase()) ||
                               atividades.notes.toLowerCase().includes(searchTerm.toLowerCase()) ||
                               atividades.intervals.some(interval => interval.notes?.toLowerCase().includes(searchTerm.toLowerCase()));
            const matchesFeeling = selectedFeeling === 'all' || atividades.feeling === selectedFeeling;
            return matchesSearch && matchesFeeling;
        })
        .sort((a, b) => {
            switch (sortBy) {
                case 'date':
                    return new Date(b.date).getTime() - new Date(a.date).getTime();
                case 'distance':
                    return b.distance - a.distance;
                case 'duration':
                    return b.duration - a.duration;
                case 'feeling':
                    const feelingOrder = ['excellent', 'good', 'regular', 'tired', 'bad'];
                    return feelingOrder.indexOf(a.feeling) - feelingOrder.indexOf(b.feeling);
                default:
                    return 0;
            }
        });

    // Calculate summary stats
    const totalDistance = atividades.reduce((sum, atividade) => sum + atividade.distance, 0);
    const totalDuration = atividades.reduce((sum, atividade) => sum + atividade.duration, 0);
    const averagePace = atividades.length > 0 ? calculatePace(totalDistance, totalDuration) : '0:00';

    const toggleActivityExpansion = (atividadeId: string) => {
        setExpandedActivities(prevExpanded => {
            const isExpanded = prevExpanded.includes(atividadeId);
            if (isExpanded) {
                return prevExpanded.filter(id => id !== atividadeId);
            } else {
                return [...prevExpanded, atividadeId];
            }
        });
    };

    // Add loading and error UI
    if (loading) {
        return (
            <div className="flex justify-center items-center h-64">
                <span className="text-slate-500 text-lg">Carregando atividades...</span>
            </div>
        );
    }
    if (error) {
        return (
            <div className="flex justify-center items-center h-64">
                <span className="text-red-500 text-lg">{error}</span>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            {/* Header */}
            <div>
                <h2 className="text-2xl font-bold text-slate-900">Histórico</h2>
                <p className="text-slate-600">Visualize e analise suas atividades</p>
            </div>

            {/* Summary Stats */}
            <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <div className="flex items-center space-x-3 mb-2">
                        <div className="p-2 rounded-lg bg-blue-100">
                            <Waves className="h-5 w-5 text-blue-600" />
                        </div>
                        <span className="text-slate-600 text-sm">Atividades Totais</span>
                    </div>
                    <div className="text-2xl font-bold text-slate-900">{atividades.length}</div>
                </div>

                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <div className="flex items-center space-x-3 mb-2">
                        <div className="p-2 rounded-lg bg-cyan-100">
                            <TrendingUp className="h-5 w-5 text-cyan-600" />
                        </div>
                        <span className="text-slate-600 text-sm">Distância Total</span>
                    </div>
                    <div className="text-2xl font-bold text-slate-900">{(totalDistance / 1000).toFixed(1)}km</div>
                </div>

                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <div className="flex items-center space-x-3 mb-2">
                        <div className="p-2 rounded-lg bg-green-100">
                            <Clock className="h-5 w-5 text-green-600" />
                        </div>
                        <span className="text-slate-600 text-sm">Tempo Total</span>
                    </div>
                    <div className="text-2xl font-bold text-slate-900">{Math.round(totalDuration / 60)}h</div>
                </div>

                <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                    <div className="flex items-center space-x-3 mb-2">
                        <div className="p-2 rounded-lg bg-purple-100">
                            <Heart className="h-5 w-5 text-purple-600" />
                        </div>
                        <span className="text-slate-600 text-sm">Ritmo Médio</span>
                    </div>
                    <div className="text-2xl font-bold text-slate-900">{averagePace}/100m</div>
                </div>
            </div>

            {/* Filters and Search */}
            <div className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200">
                <div className="flex flex-col md:flex-row gap-4">
                    {/* Search */}
                    <div className="flex-1 relative">
                        <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-slate-400" />
                        <input
                            type="text"
                            placeholder="Procure por piscina, notas, ou notas de intervalo..."
                            value={searchTerm}
                            onChange={(e) => setSearchTerm(e.target.value)}
                            className="w-full pl-10 pr-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                        />
                    </div>
                    
                    {/* Feeling Filter */}
                    <div className="flex items-center space-x-2">
                        <Filter className="h-5 w-5 text-slate-400" />
                        <select
                            value={selectedFeeling}
                            onChange={(e) => setSelectedFeeling(e.target.value)}
                            className="px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                        >
                            {feelingOptions.map(option => (
                                <option key={option.value} value={option.value}>{option.label}</option>
                            ))}
                        </select>
                    </div>
                    
                    {/* Sort */}
                    <select
                        value={sortBy}
                        onChange={(e) => setSortBy(e.target.value)}
                        className="px-4 py-3 rounded-xl border border-slate-200 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                    >
                        {sortOptions.map(option => (
                            <option key={option.value} value={option.value}>{option.label}</option>
                        ))}
                    </select>
                </div>
            </div>

            {/* Workouts List */}
            {filteredAndSortedWorkouts.length === 0 ? (
                <div className="text-center py-12">
                    <div className="text-slate-400 mb-4">
                        <Waves className="h-12 w-12 mx-auto" />
                    </div>
                    <h3 className="text-lg font-medium text-slate-900 mb-2">Não foram encontradas atividades.</h3>
                    <p className="text-slate-600">
                        {atividades.length === 0 
                            ? "Start logging your swimming workouts to see them here"
                            : "Try adjusting your search or filter criteria"
                        }
                    </p>
                </div>
            ) : (
                <div className="space-y-4">
                    {filteredAndSortedWorkouts.map((atividade) => {
                        const feelingStyle = getFeelingStyle(atividade.feeling);
                        const pace = calculatePace(atividade.distance, atividade.duration);
                        
                        return (
                            <div key={atividade.id} className="bg-white rounded-2xl p-6 shadow-sm border border-slate-200 hover:shadow-md transition-shadow">
                                <div className="flex justify-between items-start mb-4">
                                    <div className="flex items-center space-x-4">
                                        <div className="bg-gradient-to-br from-blue-400 to-cyan-500 p-3 rounded-xl">
                                            <Waves className="h-6 w-6 text-white" />
                                        </div>
                                        <div>
                                            <h3 className="text-lg font-semibold text-slate-900">{atividade.locationName}</h3>
                                            <div className="flex items-center space-x-2 text-slate-500">
                                                <Calendar className="h-4 w-4" />
                                                <span>{new Date(atividade.date).toLocaleDateString()}</span>
                                                <span className={`px-2 py-1 rounded-full text-xs font-medium ${feelingStyle.bg} ${feelingStyle.color}`}>
                                                    {atividade.feeling}
                                                </span>
                                            </div>
                                        </div>
                                    </div>
                                    <div className="text-right">
                                        <div className="text-2xl font-bold text-slate-900">{atividade.distance}m</div>
                                        <div className="text-sm text-slate-500">Ritmo: {pace}/100m</div>
                                    </div>
                                </div>
                                
                                <div className="grid grid-cols-2 md:grid-cols-5 gap-4 mb-4">
                                    <div className="flex items-center space-x-2">
                                        <Clock className="h-4 w-4 text-slate-400" />
                                        <span className="text-slate-600">{formatDuration(atividade.duration)}</span>
                                    </div>
                                    
                                    <div className="flex items-center space-x-2">
                                        <MapPin className="h-4 w-4 text-slate-400" />
                                        {atividade.locationType === 'pool' ? (
                                            <span className="text-slate-600">Piscina ({atividade.poolSize}m)</span>
                                        ) : (
                                            <span className="text-slate-600">Águas Abertas</span>
                                        )}
                                    </div>
                                    
                                    {atividade.waterTemp && (
                                        <div className="flex items-center space-x-2">
                                            <Thermometer className="h-4 w-4 text-slate-400" />
                                            <span className="text-slate-600">{atividade.waterTemp}°C</span>
                                        </div>
                                    )}
                                    
                                    {atividade.heartRateAvg && (
                                        <div className="flex items-center space-x-2">
                                            <Heart className="h-4 w-4 text-slate-400" />
                                            <span className="text-slate-600">{atividade.heartRateAvg} bpm (média)</span>
                                        </div>
                                    )}

                                    {atividade.heartRateMax && (
                                        <div className="flex items-center space-x-2">
                                            <Heart className="h-4 w-4 text-slate-400" />
                                            <span className="text-slate-600">{atividade.heartRateMax} bpm (máxima)</span>
                                        </div>
                                    )}

                                    <div className="flex items-center space-x-2">
                                        <span className="text-slate-600">{atividade.intervals.length} intervalos</span>
                                    </div>
                                </div>
                                

                                {/* Intervals List */}
                                {atividade.intervals.length > 0 && (
                                    <div className="bg-slate-50 rounded-xl p-4 mb-4">
                                        <h4 className="font-medium text-slate-900 mb-3">Intervalos</h4>
                                        <div className="space-y-2">
                                            {(() => {
                                                // Check if the current activity is expanded
                                                const isExpanded = expandedActivities.includes(atividade.id);
                                                // Determine which intervals to show
                                                const intervalsToShow = isExpanded ? atividade.intervals : atividade.intervals.slice(0, 4);
                                                return (
                                                    <>
                                                        {intervalsToShow.map((interval) => (
                                                            <div key={interval.id} className="bg-white p-3 rounded-lg border border-slate-200">
                                                                <div className="flex justify-between items-center mb-1">
                                                                    <div className="text-sm text-slate-900 font-medium">
                                                                        {interval.distance}m {interval.stroke} - {interval.time}
                                                                    </div>
                                                                </div>
                                                                {interval.notes && (
                                                                    <div className="text-xs text-slate-600 bg-blue-50 p-2 rounded border border-blue-100">
                                                                        {interval.notes}
                                                                    </div>
                                                                )}
                                                            </div>
                                                        ))}
                                                        {/* If there are more intervals than the initial slice, show the button */}
                                                        {atividade.intervals.length > 4 && (
                                                            <div className="text-center mt-3">
                                                                <button
                                                                    onClick={() => toggleActivityExpansion(atividade.id)}
                                                                    className="text-blue-600 hover:text-blue-800 text-sm font-medium"
                                                                >
                                                                    {isExpanded ? 'mostrar menos' : `mostrar mais ${atividade.intervals.length - 4} atividades`}
                                                                </button>
                                                            </div>
                                                        )}
                                                    </>
                                                );
                                            })()}
                                        </div>
                                    </div>
                                )}
                                
                                {atividade.notes && (
                                    <div className="bg-blue-50 rounded-xl p-3">
                                        <h5 className="font-medium text-slate-900 mb-1">Notas</h5>
                                        <p className="text-slate-700 text-sm">{atividade.notes}</p>
                                    </div>
                                )}
                            </div>
                        );
                    })}
                </div>
            )}
        </div>
    );
};

export default Historico;