import '../wailsjs/runtime';
import { ParseAndCalculateTime } from '../wailsjs/go/main/App';

// Global state
let timeEntries = [];

// Initialize app
document.addEventListener('DOMContentLoaded', () => {
    const pasteBtn = document.getElementById('pasteBtn');
    const clearBtn = document.getElementById('clearBtn');
    
    pasteBtn.addEventListener('click', handlePaste);
    clearBtn.addEventListener('click', handleClear);
    
    // Enable paste from keyboard
    document.addEventListener('paste', handlePasteEvent);
});

// Handle paste button click
async function handlePaste() {
    try {
        const text = await navigator.clipboard.readText();
        if (text) {
            processTimeData(text);
        }
    } catch (err) {
        console.error('Failed to read clipboard:', err);
        // Fallback for older browsers
        const textarea = document.createElement('textarea');
        textarea.style.position = 'fixed';
        textarea.style.opacity = '0';
        document.body.appendChild(textarea);
        textarea.focus();
        document.execCommand('paste');
        const text = textarea.value;
        document.body.removeChild(textarea);
        if (text) {
            processTimeData(text);
        }
    }
}

// Handle paste event (Ctrl+V)
function handlePasteEvent(event) {
    event.preventDefault();
    const text = event.clipboardData.getData('text');
    if (text) {
        processTimeData(text);
    }
}

// Process pasted time data
async function processTimeData(text) {
    // Call Go backend to parse and calculate
    const result = await ParseAndCalculateTime(text);
    
    // Add new entries to existing ones
    if (result.entries && result.entries.length > 0) {
        timeEntries = [...timeEntries, ...result.entries];
        updateDisplay();
    }
}

// Update the display
async function updateDisplay() {
    const timeList = document.getElementById('timeList');
    const totalTimeElement = document.getElementById('totalTime');
    
    // Clear current display
    timeList.innerHTML = '';
    
    // Calculate total seconds
    let totalSeconds = 0;
    
    // Display each entry
    timeEntries.forEach((entry, index) => {
        totalSeconds += entry.seconds;
        
        const entryDiv = document.createElement('div');
        entryDiv.className = 'time-entry';
        
        const timeValue = document.createElement('span');
        timeValue.className = 'time-value';
        timeValue.textContent = entry.original;
        
        const deleteBtn = document.createElement('button');
        deleteBtn.className = 'delete-btn';
        deleteBtn.innerHTML = `
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
        `;
        deleteBtn.onclick = () => handleDelete(index);
        
        entryDiv.appendChild(timeValue);
        entryDiv.appendChild(deleteBtn);
        timeList.appendChild(entryDiv);
    });
    
    // Update total time
    totalTimeElement.textContent = formatSeconds(totalSeconds);
}

// Format seconds to HH:MM:SS
function formatSeconds(totalSeconds) {
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;
    
    return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
}

// Handle delete entry
function handleDelete(index) {
    timeEntries.splice(index, 1);
    updateDisplay();
}

// Handle clear all
function handleClear() {
    timeEntries = [];
    updateDisplay();
}
