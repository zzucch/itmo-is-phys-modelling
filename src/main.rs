use full_palette::GREEN_700;
use plotters::prelude::*;

const POTENTIAL_WIDTH: f64 = 1.0;
const PERIOD: f64 = 2.0;
const BARRIER_HEIGHT: f64 = 1.0;

fn generate_linear_array(start: f64, end: f64, points_count: usize) -> Vec<f64> {
    let step = (end - start) / (points_count as f64 - 1.0);
    (0..points_count).map(|i| start + i as f64 * step).collect()
}

fn get_potential(positions: &[f64]) -> Vec<f64> {
    positions
        .iter()
        .map(|&position| {
            let n = (position / PERIOD).floor();
            if n * PERIOD + POTENTIAL_WIDTH < position && position < (n + 1.0) * PERIOD {
                BARRIER_HEIGHT
            } else {
                // n * period < x < n * period + potential_width
                0.0
            }
        })
        .collect()
}

fn compute_values(positions: &[f64], parameter: f64) -> Vec<f64> {
    positions
        .iter()
        .map(|&position| {
            if position.abs() > 1e-10 {
                parameter * (position.sin() / position) + position.cos()
            } else {
                parameter + 1.0
            }
        })
        .collect()
}

fn plot_potential(
    positions: &[f64],
    potential_values: &[f64],
) -> Result<(), Box<dyn std::error::Error>> {
    let root = BitMapBackend::new("potential.png", (1024, 768)).into_drawing_area();
    root.fill(&WHITE)?;

    let mut chart = ChartBuilder::on(&root)
        .caption(
            "potential function in kronig-penney model",
            ("sans-serif", 36).into_font(),
        )
        .margin(10)
        .x_label_area_size(30)
        .y_label_area_size(30)
        .build_cartesian_2d(positions[0]..positions[positions.len() - 1], -0.5..1.5)?;

    chart
        .configure_mesh()
        .x_desc("position x")
        .y_desc("potential U(x)")
        .draw()?;
    chart.draw_series(LineSeries::new(
        positions
            .iter()
            .copied()
            .zip(potential_values.iter().copied()),
        &RED,
    ))?;
    Ok(())
}

fn plot_graphical_analysis(
    positions: &[f64],
    model_values: &[f64],
    parameter: f64,
) -> Result<(), Box<dyn std::error::Error>> {
    let filename = &format!("{parameter:.5}.png");

    let root = BitMapBackend::new(filename, (1024, 768)).into_drawing_area();
    root.fill(&WHITE)?;

    let mut chart = ChartBuilder::on(&root)
        .caption("kronig-penney model", ("sans-serif", 36).into_font())
        .margin(10)
        .x_label_area_size(30)
        .y_label_area_size(30)
        .build_cartesian_2d(positions[0]..positions[positions.len() - 1], -1.5..2.5)?;

    chart.configure_mesh().x_desc("a * alpha").draw()?;
    chart.draw_series(LineSeries::new(
        positions.iter().copied().zip(model_values.iter().copied()),
        &BLUE,
    ))?;
    chart.draw_series(LineSeries::new(
        positions.iter().map(|&xi| (xi, 1.0)),
        &GREEN_700,
    ))?;
    chart.draw_series(LineSeries::new(
        positions.iter().map(|&xi| (xi, -1.0)),
        &GREEN_700,
    ))?;
    Ok(())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let x = generate_linear_array(-10.0, 11.0, 1000);

    let v = get_potential(&x);
    plot_potential(&x, &v)?;

    let p = 1.0;
    let y = compute_values(&x, p);
    plot_graphical_analysis(&x, &y, p)?;

    Ok(())
}
